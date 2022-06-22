package client

import (
	"github.com/juxuny/yc/services"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"log"
	"os"
	"path"
	"strings"
)

func (t *GenCommand) genGo() {
	serviceName := t.getServiceName()
	log.Println("generating go sdk:", serviceName)
	service := services.NewServiceEntity(serviceName)
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
	for _, item := range result.ProtoBody {
		switch item.(type) {
		case *parser.Package:
		case *parser.Option:
		case *parser.Service:
			svc = append(svc, item.(*parser.Service))
		case *parser.Message:
		}
	}
	if len(svc) > 1 {
		log.Fatal("only support one service in each *.proto")
	}
	t.genGoService(service, svc[0])
}

func (t *GenCommand) genGoService(serviceEntity services.ServiceEntity, svc *parser.Service) {
	moduleName, err := utils.GetCurrentPackageName(t.WorkDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("module name:", moduleName)
	handlerEntity := services.HandlerInitEntity{
		ServiceEntity: serviceEntity,
		GoModuleName:  moduleName,
	}
	methodsMap := make(map[string][]services.MethodEntity)
	for _, m := range svc.ServiceBody {
		rpc := m.(*parser.RPC)
		groupName := services.GetContentByProtoTagFistOne(services.ProtoTagGroup, rpc.Comments)
		if groupName == "" {
			groupName = "default"
		}
		if !t.Internal && services.CheckIfContainProtoTag(services.ProtoTagInternal, rpc.Comments) {
			log.Println("ignore internal rpc:", rpc.RPCName)
			continue
		}
		method := services.MethodEntity{
			HandlerInitEntity: handlerEntity,
			Desc:              services.GetDescFromFieldCommentsOfProto(rpc.Comments),
			Group:             groupName,
			Api:               serviceEntity.ServiceName + "/" + strings.ReplaceAll(utils.ToUnderLine(rpc.RPCName), "_", "-"),
			MethodName:        rpc.RPCName,
			Request:           rpc.RPCRequest.MessageType,
			Response:          rpc.RPCResponse.MessageType,
			UseAuth:           services.GetAuthFromRpcCommentsOfProto(rpc.Comments),
		}
		methodsMap[groupName] = append(methodsMap[groupName], method)
	}
	groups := make([]string, 0)
	for g := range methodsMap {
		groups = append(groups, g)
	}
	utils.StringHelper.SortSlice(groups)
	var clientSdkEntity = services.ClientSdkEntity{
		Methods:       make([]services.MethodEntity, 0),
		ServiceEntity: serviceEntity,
	}
	for _, g := range groups {
		clientSdkEntity.Methods = append(clientSdkEntity.Methods, methodsMap[g]...)
	}
	outputFile := path.Join(t.WorkDir, "client_rpc.sdk.go")
	if err := template.RunEmbedFile(templateFs, goClientRpcSdkFileName, outputFile, clientSdkEntity); err != nil {
		log.Fatal(err)
	}
	clientSdkInitFile := path.Join(t.WorkDir, "client.sdk.go")
	if err := template.RunEmbedFile(templateFs, goClientSdkFileName, clientSdkInitFile, clientSdkEntity); err != nil {
		log.Fatal(err)
	}
}
