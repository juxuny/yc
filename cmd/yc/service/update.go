package service

import (
	"fmt"
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/services"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"github.com/spf13/cobra"
	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type UpdateCommand struct {
	WorkDir string
}

func (t *UpdateCommand) BeforeRun(cmd *cobra.Command) {
	if Env.Gopath == "" {
		log.Fatal("missing environment GOPATH")
	}
	if Env.YcHome == "" {
		log.Fatal("missing environment YC_HOME")
	}
}

func (t *UpdateCommand) Prepare(cmd *cobra.Command) {
}

func (t *UpdateCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.WorkDir, "work-dir", "w", "", "working dir")
}

func (t *UpdateCommand) Run() {
	if t.WorkDir == "" {
		if w, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			t.WorkDir = w
		}
	}

	serviceName := t.getServiceName()
	if serviceName == "" {
		log.Fatal("not found proto file in working directory")
	}
	log.Println("service name: ", serviceName)
	service := services.NewServiceEntity(serviceName)

	// init env config
	if err := t.initEnvConfig(service); err != nil {
		log.Fatal(err)
	}

	// auto generate client and request entities
	t.genRpc(service)
	t.genExtend(service)
}

func (t *UpdateCommand) getServiceName() string {
	serviceName := ""
	if err := filepath.Walk(t.WorkDir, func(filePath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if path.Ext(filePath) == ".proto" {
			serviceName = path.Base(filePath)
			return nil
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return utils.StringHelper.TrimSubSequenceRight(serviceName, ".proto")
}

func (t *UpdateCommand) initEnvConfig(service services.ServiceEntity) error {
	configDir := path.Join(t.WorkDir, "config")
	if err := utils.TouchDir(configDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := template.RunEmbedFile(templateFs, envConfigFileName, path.Join(t.WorkDir, "config", "env.go"), service); err != nil {
		log.Fatal("create env.go failed:", err)
	}
	return nil
}

func (t *UpdateCommand) genRpc(service services.ServiceEntity) {
	args := []string{
		"-I=" + Env.YcHome,
		"-I=.",
		"--go_out=.",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=.",
		"--go-grpc_opt=paths=source_relative",
		fmt.Sprintf("%s.proto", service.ProtoFileName),
	}
	if err := cmd.Exec(
		"protoc",
		args...,
	); err != nil {
		log.Fatal(err)
	}
}

func (t *UpdateCommand) genExtend(service services.ServiceEntity) {
	reader, err := os.Open(path.Join(t.WorkDir, service.ProtoFileName+".proto"))
	if err != nil {
		log.Fatal("parse proto failed: ", err)
	}
	defer reader.Close()
	result, err := protoparser.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	messages := make([]*parser.Message, 0)
	svc := make([]*parser.Service, 0)
	for _, item := range result.ProtoBody {
		switch item.(type) {
		case *parser.Package:
		case *parser.Option:
		case *parser.Service:
			svc = append(svc, item.(*parser.Service))
		case *parser.Message:
			messages = append(messages, item.(*parser.Message))
		}
	}
	t.genValidator(service, messages)
	t.genService(service, svc)
}

func (t *UpdateCommand) genService(service services.ServiceEntity, svc []*parser.Service) {
	t.genHandler(service, svc)
	for _, s := range svc {
		log.Println(s.ServiceName)
	}
}

func (t *UpdateCommand) genHandler(service services.ServiceEntity, svc []*parser.Service) {
	moduleName, err := utils.GetCurrentPackageName(t.WorkDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("module name:", moduleName)
	// create default
	if err := utils.TouchDir(path.Join(t.WorkDir, "handler"), 0755); err != nil {
		log.Fatal(err)
	}
	handlerEntity := services.HandlerEntity{
		ServiceEntity: service,
		GoModuleName:  moduleName,
	}
	if err := template.RunEmbedFile(templateFs, handlerDefaultFileName, path.Join(t.WorkDir, "handler", "default.go"), handlerEntity); err != nil {
		log.Fatal(err)
	}
	if err := template.RunEmbedFile(templateFs, handlerWrapperFileName, path.Join(t.WorkDir, "handler", "wrapper.go"), handlerEntity); err != nil {
		log.Fatal(err)
	}
	// create method groups
}

func (t *UpdateCommand) genValidator(service services.ServiceEntity, msgs []*parser.Message) {
	validatorEntities := services.ValidatorEntity{
		ServiceEntity: service,
	}
	for _, m := range msgs {
		//log.Println(m.MessageName)
		fields := make([]services.MessageField, 0)
		if !strings.HasSuffix(m.MessageName, "Request") && !strings.HasSuffix(m.MessageName, "Req") {
			continue
		}
		for _, body := range m.MessageBody {
			//log.Println(reflect.ValueOf(body).Type().String())
			f, ok := body.(*parser.Field)
			if !ok {
				log.Println("is not a message field")
				continue
			}
			errorMessageTemplate := ""
			formulas := make([]services.ValidatorFormula, 0)
			for _, comment := range f.Comments {
				if len(formulas) > 0 && errorMessageTemplate != "" {
					break
				}
				for _, l := range comment.Lines() {
					l = strings.TrimSpace(l)
					if strings.Index(l, "@v:") == 0 {
						kvs := strings.Split(strings.TrimSpace(utils.StringHelper.TrimSubSequenceLeft(l, "@v:")), "=")
						if len(kvs) >= 2 {
							formulas = append(formulas, services.ValidatorFormula{
								Pattern:  strings.TrimSpace(kvs[0]),
								RefValue: strings.TrimSpace(kvs[1]),
							})
						}
						continue
					}
					if strings.Index(l, "@msg:") == 0 {
						errorMessageTemplate = strings.TrimSpace(utils.StringHelper.TrimSubSequenceLeft(l, "@msg:"))
						continue
					}
				}
			}
			if len(formulas) > 0 && errorMessageTemplate != "" {
				fields = append(fields, services.MessageField{
					Name:      utils.ToUpperFirst(utils.ToHump(f.FieldName)),
					Formulas:  formulas,
					ParamName: utils.ToLowerFirst(utils.ToHump(f.FieldName)),
					Error:     errorMessageTemplate,
				})
			}
		}
		messageItem := services.Message{
			Name:   m.MessageName,
			Fields: fields,
		}
		validatorEntities.Messages = append(validatorEntities.Messages, messageItem)
	}

	if err := template.RunEmbedFile(templateFs, validatorFileName, path.Join(t.WorkDir, service.ProtoFileName+".pb.validator.go"), validatorEntities); err != nil {
		log.Fatal(err)
	}
}

func init() {
	builder := cmd.NewCommandBuilder("update", &UpdateCommand{})
	Command.AddCommand(builder.Build())
}
