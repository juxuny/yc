package client

import (
	"github.com/juxuny/yc"
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

func (t *GenCommand) genReactTs() {
	serviceName := t.getServiceName()
	log.Println("generating react-ts sdk:", serviceName)
	service := services.NewServiceEntity(serviceName, yc.Version)
	reader, err := os.Open(path.Join(t.WorkDir, service.ProtoFileName+".proto"))
	if err != nil {
		log.Fatal("parse proto failed: ", err)
	}
	defer reader.Close()
	result, err := protoparser.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	svc := make([]*parser.Service, 0)
	messages := make([]*parser.Message, 0)
	enums := make([]*parser.Enum, 0)
	methods := make([]*parser.RPC, 0)
	for _, item := range result.ProtoBody {
		switch item.(type) {
		case *parser.Package:
		case *parser.Option:
		case *parser.Service:
			svc = append(svc, item.(*parser.Service))
		case *parser.Message:
			messages = append(messages, item.(*parser.Message))
		case *parser.Enum:
			enums = append(enums, item.(*parser.Enum))
		}
	}
	if len(svc) > 1 {
		log.Fatal("only support one service in each *.proto")
	}
	t.genReactTsIndex(service)
	t.genReactTsTyping(service, messages, enums)

	for _, item := range svc[0].ServiceBody {
		switch item.(type) {
		case *parser.RPC:
			methods = append(methods, item.(*parser.RPC))
		}
	}
	t.genReactTsMethods(service, methods)
}

func (t *GenCommand) genReactTsIndex(service services.ServiceEntity) {
	outputFile := path.Join(t.ReactTs, "index.ts")
	log.Println("gen index.ts: ", outputFile)
	if err := template.RunEmbedFile(templateFs, reactTsIndexFileName, outputFile, service); err != nil {
		log.Fatal(err)
	}
}

func (t *GenCommand) getReactTsEnumFields(enumName string, enum *parser.Enum) []services.EnumField {
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

func (t *GenCommand) getReactTsEnumValueSet(enum *parser.Enum) string {
	values := make([]string, 0)
	for _, item := range enum.EnumBody {
		switch item.(type) {
		case *parser.EnumField:
			field := item.(*parser.EnumField)
			values = append(values, field.Number)
		}
	}
	return strings.Join(values, " | ")
}

func (t *GenCommand) getReactTsMessageFields(message *parser.Message) []services.ReactTsMessageField {
	ret := make([]services.ReactTsMessageField, 0)
	for _, m := range message.MessageBody {
		switch m.(type) {
		case *parser.Field:
			field := m.(*parser.Field)
			finalDataType, nullable := services.ConvertProtoTypeToReactTsDataType(field.Type)
			if field.IsRepeated {
				finalDataType = finalDataType + "[]"
			}
			ret = append(ret, services.ReactTsMessageField{
				Name:     field.FieldName,
				Desc:     services.GetDescFromFieldCommentsOfProto(field.Comments),
				Required: services.CheckIfRequired(field.Comments) && !nullable,
				Type:     finalDataType,
			})
		}
	}
	return ret
}

func (t *GenCommand) genReactTsTyping(service services.ServiceEntity, messages []*parser.Message, enums []*parser.Enum) {
	outputFile := path.Join(t.ReactTs, "typing.d.ts")
	log.Println("gen typing.d.ts:", outputFile)
	reactTsTypingEntity := services.ReactTsTypingEntity{
		ServiceEntity: service,
		Enums:         make([]services.EnumEntity, 0),
		Messages:      make([]services.ReactTsMessage, 0),
	}
	for _, item := range enums {
		reactTsTypingEntity.Enums = append(reactTsTypingEntity.Enums, services.EnumEntity{
			Desc:     services.GetDescFromFieldCommentsOfProto(item.Comments),
			EnumName: item.EnumName,
			Fields:   t.getReactTsEnumFields(item.EnumName, item),
			ValueSet: t.getReactTsEnumValueSet(item),
		})
	}
	for _, item := range messages {
		if strings.Index(item.MessageName, "Model") == 0 {
			continue
		}
		reactTsTypingEntity.Messages = append(reactTsTypingEntity.Messages, services.ReactTsMessage{
			Name:   item.MessageName,
			Fields: t.getReactTsMessageFields(item),
			Desc:   services.GetDescFromFieldCommentsOfProto(item.Comments),
		})
	}
	if err := template.RunEmbedFile(templateFs, reactTsTypingFileName, outputFile, reactTsTypingEntity); err != nil {
		log.Fatal(err)
	}
}

func (t *GenCommand) genReactTsMethods(service services.ServiceEntity, methods []*parser.RPC) {
	apiEntity := services.ReactTsApiEntity{
		ServiceEntity: service,
		Methods:       []services.ReactTsMethod{},
	}
	for _, item := range methods {
		if !t.Internal && services.CheckIfContainProtoTag(services.ProtoTagInternal, item.Comments) {
			log.Println("ignore internal rpc: ", item.RPCName)
			continue
		}
		apiEntity.Methods = append(apiEntity.Methods, services.ReactTsMethod{
			ServiceEntity: service,
			Api:           strings.ReplaceAll(service.ServiceName, "_", "-") + "/" + strings.ReplaceAll(utils.ToUnderLine(item.RPCName), "_", "-"),
			MethodName:    item.RPCName,
			Request:       item.RPCRequest.MessageType,
			Response:      item.RPCResponse.MessageType,
			Desc:          services.GetDescFromFieldCommentsOfProto(item.Comments),
		})
	}
	output := path.Join(t.ReactTs, "api.ts")
	if err := template.RunEmbedFile(templateFs, reactTsApiFilename, output, apiEntity); err != nil {
		log.Fatal(err)
	}
}
