package service

import (
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/services"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"github.com/spf13/cobra"
	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type UpdateCommand struct {
	WorkDir  string
	Override bool
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
	cmd.PersistentFlags().BoolVar(&t.Override, "override", false, "override all auto-gen file")
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

	// create service entity from *.proto file in current directory
	service := t.getServiceEntity()

	// init env config
	if err := t.initEnvConfig(service); err != nil {
		log.Fatal(err)
	}

	// auto generate client and request entities
	t.genRpc(service)
	t.genExtend(service)
	t.fmt()
}

func (t *UpdateCommand) getServiceEntity() services.ServiceEntity {
	serviceName := t.getServiceName()
	if serviceName == "" {
		log.Fatal("get service name failed")
	}
	log.Println("service name: ", serviceName)
	ret := services.NewServiceEntity(serviceName, yc.Version)
	reader, err := os.Open(path.Join(t.WorkDir, ret.ProtoFileName+".proto"))
	if err != nil {
		log.Fatal("parse proto failed: ", err)
	}
	defer reader.Close()
	result, err := protoparser.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	var service *parser.Service
	for _, m := range result.ProtoBody {
		if service != nil {
			break
		}
		switch m.(type) {
		case *parser.Service:
			service = m.(*parser.Service)
			break
		}
	}
	if service == nil {
		log.Fatal("not found service definition")
	}
	serviceVersion := services.GetContentByProtoTagFistOne(services.ProtoTagVersion, service.Comments)
	if serviceVersion == "" {
		log.Fatal("service missing @version")
	}
	serviceLevel := services.GetContentByProtoTagFistOne(services.ProtoTagLevel, service.Comments)
	if serviceLevel == "" {
		log.Fatal("service missing @level")
	}
	if _, err := strconv.ParseInt(serviceLevel, 10, 64); err != nil {
		log.Fatalf("parse level failed, is not a integer: %v", err)
	}
	serviceCode := services.GetContentByProtoTagFistOne(services.ProtoTagCode, service.Comments)
	if serviceCode == "" {
		log.Fatalf("@service missing @code")
	}
	if _, err := strconv.ParseInt(serviceCode, 10, 64); err != nil {
		log.Fatalf("parse service code failed, is not a integer: %v", serviceCode)
	}
	ret.Version = serviceVersion
	ret.Level = serviceLevel
	ret.ServiceCode = serviceCode
	return ret
}

func (t *UpdateCommand) getServiceName() string {
	serviceName := ""
	if _, err := os.Stat(t.WorkDir); os.IsNotExist(err) {
		log.Fatal(err)
	}
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
	outEnvFile := path.Join(t.WorkDir, "config", "env.go")
	if _, err := os.Stat(outEnvFile); os.IsNotExist(err) {
		if err := template.RunEmbedFile(templateFs, envConfigFileName, outEnvFile, service); err != nil {
			log.Fatal("create env.go failed:", err)
		}
	}
	return nil
}

func (t *UpdateCommand) genRpc(service services.ServiceEntity) {
	args := []string{
		"-I=" + Env.YcHome,
		"-I=.",
		"--go_out=.",
		"--go_opt=paths=source_relative",
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
	t.genExt(service, messages)
	t.genService(service, svc)
}

func (t *UpdateCommand) getIgnoreAuthMethodList(serviceList []*parser.Service) []string {
	ret := make([]string, 0)
	for _, s := range serviceList {
		for _, item := range s.ServiceBody {
			rpc, ok := item.(*parser.RPC)
			if !ok {
				continue
			}
			if services.CheckIfContainProtoTag(services.ProtoTagIgnoreAuth, rpc.Comments) {
				ret = append(ret, strings.ReplaceAll(utils.StringHelper.ToUnderLine(rpc.RPCName), "_", "-"))
			}
		}
	}
	return ret
}

func (t *UpdateCommand) getOpenApiList(serviceList []*parser.Service) []string {
	ret := make([]string, 0)
	for _, s := range serviceList {
		for _, item := range s.ServiceBody {
			rpc, ok := item.(*parser.RPC)
			if !ok {
				continue
			}
			if services.CheckIfContainProtoTag(services.ProtoTagCheckSign, rpc.Comments) {
				ret = append(ret, strings.ReplaceAll(utils.StringHelper.ToUnderLine(rpc.RPCName), "_", "-"))
			}
		}
	}
	return ret
}

func (t *UpdateCommand) genService(service services.ServiceEntity, svc []*parser.Service) {
	t.genHandler(service, svc)
	for _, s := range svc {
		log.Println(s.ServiceName)
	}

	// gen entrypoint
	if _, err := os.Stat(path.Join(t.WorkDir, "default.go")); t.Override || os.IsNotExist(err) {
		if err := template.RunEmbedFile(templateFs, defaultConfigFileName, path.Join(t.WorkDir, "default.go"), service); err != nil {
			log.Fatal("create default config failed:", err)
		}
	}
	if err := utils.TouchDirs([]string{
		path.Join(t.WorkDir, "server"),
		path.Join(t.WorkDir, "server", "http"),
	}, 0755); err != nil {
		log.Fatal(err)
	}
	moduleName, err := utils.GetCurrentPackageName(t.WorkDir)
	if err != nil {
		log.Fatal(err)
	}
	if err := template.RunEmbedFile(templateFs, httpServerFileName, path.Join(t.WorkDir, "server", "http", "http_server.go"), services.EntrypointEntity{
		ServiceEntity:  service,
		GoModuleName:   moduleName,
		OpenApiList:    t.getOpenApiList(svc),
		IgnoreAuthList: t.getIgnoreAuthMethodList(svc),
	}); err != nil {
		log.Fatal(err)
	}

	// init main.go
	log.Println("create main file")
	outMainFile := path.Join(t.WorkDir, "server", "main.go")
	if _, err := os.Stat(outMainFile); os.IsNotExist(err) {
		if err := template.RunEmbedFile(templateFs, mainFileName, outMainFile, services.EntrypointEntity{
			ServiceEntity: service,
			GoModuleName:  moduleName,
		}); err != nil {
			log.Fatal("create main file error:", err)
		}
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
	handlerEntity := services.HandlerInitEntity{
		ServiceEntity: service,
		GoModuleName:  moduleName,
	}
	handlerDefaultOutputFile := path.Join(t.WorkDir, "handler", "default.go")
	if _, err := os.Stat(handlerDefaultOutputFile); os.IsNotExist(err) {
		if err := template.RunEmbedFile(templateFs, handlerDefaultFileName, path.Join(t.WorkDir, "handler", "default.go"), handlerEntity); err != nil {
			log.Fatal(err)
		}
	}
	//handlerWrapperOutputFile := path.Join(t.WorkDir, "handler", "wrapper.go")
	//if _, err := os.Stat(handlerWrapperOutputFile); os.IsNotExist(err) {
	//	if err := template.RunEmbedFile(templateFs, handlerWrapperFileName, path.Join(t.WorkDir, "handler", "wrapper.go"), handlerEntity); err != nil {
	//		log.Fatal(err)
	//	}
	//}
	// create method groups
	if len(svc) > 1 {
		log.Fatal("not allow multiple services in one proto file")
	}
	if len(svc) == 0 {
		log.Println("service definition not found")
		return
	}
	methodsMap := make(map[string][]services.MethodEntity)
	for _, m := range svc[0].ServiceBody {
		rpc := m.(*parser.RPC)
		groupName, ok := services.GetGroupNameFromRpcCommentsOfProto(rpc.Comments)
		if !ok {
			groupName = "default"
		}
		method := services.MethodEntity{
			HandlerInitEntity: handlerEntity,
			MethodName:        rpc.RPCName,
			Request:           rpc.RPCRequest.MessageType,
			Response:          rpc.RPCResponse.MessageType,
			UseAuth:           services.GetAuthFromRpcCommentsOfProto(rpc.Comments),
		}
		methodsMap[groupName] = append(methodsMap[groupName], method)
	}
	for group, methods := range methodsMap {
		t.checkOrInitMethodGroupFile(group, handlerEntity)
		for _, m := range methods {
			t.checkOrAppendMethod(group, m)
		}
	}
}

func (t *UpdateCommand) checkOrInitMethodGroupFile(group string, handlerInitEntity services.HandlerInitEntity) {
	methodGroupFileName := t.getMethodFileNameFromGroup(group)
	stat, err := os.Stat(path.Join(t.WorkDir, "handler", methodGroupFileName))
	if err == nil {
		if stat.IsDir() {
			log.Fatalf("'%s' cannot be a directory", methodGroupFileName)
		}
		return
	}
	if err := template.RunEmbedFile(templateFs, handlerMethodInitFileName, path.Join(t.WorkDir, "handler", methodGroupFileName), handlerInitEntity); err != nil {
		log.Fatal("init method group file failed:", err)
	}
}
func (t *UpdateCommand) checkOrAppendMethod(group string, method services.MethodEntity) {
	methodGroupFileName := t.getMethodFileNameFromGroup(group)
	if found, err := t.checkIfContainMethod(path.Join(t.WorkDir, "handler", methodGroupFileName), group, method); err != nil {
		log.Fatalf("create method %s failed: %v", method.MethodName, err)
	} else if !found {
		if err := template.AppendFromEmbedFile(templateFs, handlerMethodFuncFileName, path.Join(t.WorkDir, "handler", methodGroupFileName), method); err != nil {
			log.Fatal("create rpc method failed:", err)
		}
	}
}

func (t *UpdateCommand) checkIfContainMethod(fileName string, group string, method services.MethodEntity) (found bool, err error) {
	searchKey := fmt.Sprintf("func %s(ctx context.Context, req *%s.%s) (resp *%s.%s, err error)", method.MethodName, method.PackageAlias, method.Request, method.PackageAlias, method.Response)
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false, err
	}
	keys := strings.Split(searchKey, " ")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if utils.StringHelper.ContainsAllKey(line, keys) {
			return true, nil
		}
	}
	return
}

func (t *UpdateCommand) getMethodFileNameFromGroup(group string) string {
	return fmt.Sprintf("method_%s.go", utils.ToUnderLine(group))
}

func (t *UpdateCommand) getWrapperFileNameFromGroup(group string) string {
	return fmt.Sprintf("wrapper_%s.go", utils.ToUnderLine(group))
}

func (t *UpdateCommand) genValidator(service services.ServiceEntity, msgs []*parser.Message) {
	validatorEntities := services.ValidatorEntity{
		ServiceEntity: service,
	}
	for _, m := range msgs {
		//log.Println(m.MessageName)
		fields := make([]services.ValidatorMessageField, 0)
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
						l = strings.TrimSpace(utils.StringHelper.TrimSubSequenceLeft(l, "@v:"))
						stopIndex := strings.Index(l, "=")
						if stopIndex > 0 {
							pattern := l[:stopIndex]
							refValue := l[(stopIndex + 1):]
							formulas = append(formulas, services.ValidatorFormula{
								Pattern:  strings.TrimSpace(pattern),
								RefValue: strings.TrimSpace(refValue),
							})
						} else {
							formulas = append(formulas, services.ValidatorFormula{
								Pattern:  strings.TrimSpace(l),
								RefValue: "",
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
			if len(formulas) > 0 {
				fields = append(fields, services.ValidatorMessageField{
					Name:      utils.ToUpperFirst(utils.ToHump(f.FieldName)),
					Formulas:  formulas,
					ParamName: utils.ToLowerFirst(utils.ToHump(f.FieldName)),
					Error:     errorMessageTemplate,
				})
			}
		}
		messageItem := services.ValidatorMessage{
			Name:   m.MessageName,
			Fields: fields,
		}
		validatorEntities.Messages = append(validatorEntities.Messages, messageItem)
	}

	if err := template.RunEmbedFile(templateFs, validatorFileName, path.Join(t.WorkDir, service.ProtoFileName+".pb.validator.go"), validatorEntities); err != nil {
		log.Fatal(err)
	}
}

func (t *UpdateCommand) genExt(service services.ServiceEntity, messages []*parser.Message) {
	cloneEntity := services.CloneEntity{
		ServiceEntity: service,
		Messages:      []services.CloneMessage{},
	}
	for _, m := range messages {
		cloneEntity.Messages = append(cloneEntity.Messages, services.CloneMessage{
			Name: m.MessageName,
		})
	}
	output := path.Join(t.WorkDir, service.ProtoFileName+"_ext.pb.go")
	if err := template.RunEmbedFile(templateFs, extFileName, output, cloneEntity); err != nil {
		log.Fatal(err)
	}
}

func (t *UpdateCommand) fmt() {
	args := []string{
		"fmt", "./...",
	}
	if err := cmd.Exec(
		"go",
		args...,
	); err != nil {
		log.Fatal(err)
	}
}

func init() {
	builder := cmd.NewCommandBuilder("update", &UpdateCommand{})
	Command.AddCommand(builder.Build())
}
