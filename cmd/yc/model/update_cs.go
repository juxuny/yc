package model

import (
	"github.com/juxuny/yc/services"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
)

// genCsModel generate c-sharper model file
func (t *UpdateCommand) genCsModel(service services.ServiceEntity) {
	protoFile := t.getProtoFileName(service)
	log.Println("parsing proto file: ", protoFile)
	reader, err := os.Open(protoFile)
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
		switch it := item.(type) {
		case *parser.Enum:
			internalDataType[it.EnumName] = true
		case *parser.Message:
			internalDataType[it.MessageName] = true
			messages = append(messages, it)
		}
	}
	refMap := make(map[string][]services.RefModel)
	for _, m := range messages {
		refModel, found := getRefFromMessageOfProto(m)
		if found {
			refMap[refModel] = append(refMap[refModel], t.createRefModelFromMessageOfProto(service, m, internalDataType))
		}
	}
	for _, m := range messages {
		if strings.Index(m.MessageName, "Model") == 0 {
			t.createCSharperModel(service, m, internalDataType, refMap)
		}
	}
}

func (t *UpdateCommand) genCsEnum(service services.ServiceEntity) {
	protoFile := t.getProtoFileName(service)
	log.Println("parsing proto file: ", protoFile)
	reader, err := os.Open(protoFile)
	if err != nil {
		log.Fatal("parse proto failed: ", err)
	}
	defer reader.Close()
	result, err := protoparser.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	enums := make([]*parser.Enum, 0)
	for _, item := range result.ProtoBody {
		switch it := item.(type) {
		case *parser.Enum:
			enums = append(enums, it)
		}
	}
	for _, m := range enums {
		if strings.Index(m.EnumName, "Enum") == 0 {
			t.createCSharperEnum(service, m)
		}
	}
}

func (t *UpdateCommand) getCsEnumFields(enumName string, enum *parser.Enum) []services.EnumField {
	ret := make([]services.EnumField, 0)
	for _, item := range enum.EnumBody {
		switch item.(type) {
		case *parser.EnumField:
			enumField := item.(*parser.EnumField)
			ret = append(ret, services.EnumField{
				FieldName: strings.Replace(enumField.Ident, enumName, "", 1),
				Value:     enumField.Number,
				Desc:      services.GetDescFromFieldCommentsOfProto(enumField.Comments),
			})
		default:
			log.Println("unknown type:", reflect.TypeOf(item).String())
		}
	}
	return ret
}

func (t *UpdateCommand) createCSharperEnum(service services.ServiceEntity, enum *parser.Enum) {
	log.Println("create enum:", utils.ToUnderLine(enum.EnumName))
	outEnumFile := path.Join(t.ModelOutputDir, enum.EnumName+".cs")
	enumEntity := services.EnumEntity{
		ServiceEntity:   service,
		EnumName:        enum.EnumName,
		Fields:          t.getCsEnumFields(enum.EnumName, enum),
		Desc:            services.GetDescFromFieldCommentsOfProto(enum.Comments),
		CSharpNamespace: t.CSharpModelNamespace,
	}
	if err := template.RunEmbedFile(templateFs, csEnumFileName, outEnumFile, enumEntity); err != nil {
		log.Fatal(err)
	}
}

func (t *UpdateCommand) createCSharperModel(service services.ServiceEntity, msg *parser.Message, internalDataType map[string]bool, refMap map[string][]services.RefModel) {
	log.Println("create model:", utils.ToUnderLine(msg.MessageName))
	outModelFile := path.Join(t.ModelOutputDir, msg.MessageName+".cs")
	model := services.ModelEntity{
		ServiceEntity:        service,
		CSharpModelNamespace: t.CSharpModelNamespace,
		CSharpBaseNamespace:  t.CSharpBaseNamespace,
		Model:                t.createModelFromMessageOfProto(service, msg, internalDataType, refMap[msg.MessageName]),
	}
	if err := template.RunEmbedFile(templateFs, csModelFileName, outModelFile, model); err != nil {
		log.Fatal(err)
	}
}
