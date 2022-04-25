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
)

type UpdateCommand struct {
	WorkDir string
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
	t.genEntrypoint(service)
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
	if err := cmd.Exec("protoc", "--go_out=.", "--go_opt=paths=source_relative", "--go-grpc_out=.", "--go-grpc_opt=paths=source_relative", fmt.Sprintf("%s.proto", service.ProtoFileName)); err != nil {
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
}

func (t *UpdateCommand) genValidator(service services.ServiceEntity, msgs []*parser.Message) {

}

func (t *UpdateCommand) genEntrypoint(service services.ServiceEntity) {

}

func init() {
	builder := cmd.NewCommandBuilder("update", &UpdateCommand{})
	Command.AddCommand(builder.Build())
}
