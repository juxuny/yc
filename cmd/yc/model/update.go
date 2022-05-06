package model

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

func (t *UpdateCommand) Prepare(cmd *cobra.Command) {
	log.Println("prepare")
}

func (t *UpdateCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.WorkDir, "work-dir", "w", "", "working dir")
}

func (t *UpdateCommand) BeforeRun(cmd *cobra.Command) {
	log.Println("before")
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

func (t *UpdateCommand) Run() {
	if t.WorkDir == "" {
		if w, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			t.WorkDir = w
		}
	}
	log.Println("getting service name...")
	serviceName := t.getServiceName()
	if serviceName == "" {
		log.Fatal("not found proto file in working directory")
	}
	log.Println("service name: ", serviceName)
	service := services.NewServiceEntity(serviceName)
	t.genRpc(service)
	t.genModel(service)
	t.fmt()
}

func (t *UpdateCommand) genModel(service services.ServiceEntity) {
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
	internalDataType := make(map[string]bool)
	for _, item := range result.ProtoBody {
		switch item.(type) {
		case *parser.Enum:
			internalDataType[item.(*parser.Enum).EnumName] = true
		case *parser.Message:
			internalDataType[item.(*parser.Message).MessageName] = true
			messages = append(messages, item.(*parser.Message))
		}
	}
	for _, m := range messages {
		if strings.Index(m.MessageName, "Model") == 0 {
			t.createModel(service, m, internalDataType)
		}
	}
}

func (t *UpdateCommand) createModel(service services.ServiceEntity, msg *parser.Message, internalDataType map[string]bool) {
	log.Println("create model:", utils.ToUnderLine(msg.MessageName))
	outModelFile := path.Join(t.WorkDir, "db", utils.ToUnderLine(msg.MessageName)+".go")
	moduleName, err := utils.GetCurrentPackageName(t.WorkDir)
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]services.ModelField, 0)
	hasDeletedAt := findDeletedAtFromMessageOfProto(msg)
	for _, item := range msg.MessageBody {
		f, ok := item.(*parser.Field)
		if !ok {
			log.Fatal("is not a message field")
		}
		field := services.ModelField{
			TableName:     strings.Replace(msg.MessageName, "Model", "Table", 1),
			ModelName:     msg.MessageName,
			FieldName:     f.FieldName,
			OrmFieldName:  utils.ToUnderLine(f.FieldName),
			ModelDataType: f.Type,
			HasDeletedAt:  hasDeletedAt,
			Ignore:        getIgnoreFromMessageCommentsOfProto(f.Comments),
			HasIndex:      getIndexFromMessageCommentsOfProto(f.Comments),
		}
		if internalDataType[f.Type] {
			field.ModelDataType = strings.Join([]string{service.PackageAlias, f.Type}, ".")
		}
		fields = append(fields, field)
	}
	model := services.ModelEntity{
		ServiceEntity: service,
		GoModuleName:  moduleName,
		Model: services.Model{
			TableNameWithoutServicePrefix: utils.ToUnderLine(strings.Trim(strings.Replace(msg.MessageName, "Model", "", 1), "_")),
			Fields:                        fields,
			TableName:                     strings.Replace(msg.MessageName, "Model", "Table", 1),
			ModelName:                     msg.MessageName,
			HasDeletedAt:                  hasDeletedAt,
		},
	}
	if err := template.RunEmbedFile(templateFs, modelFileName, outModelFile, model); err != nil {
		log.Fatal(err)
	}
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
