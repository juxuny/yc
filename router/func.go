package router

import (
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

func CreateMetadataFromHeader(header http.Header) metadata.MD {
	ret := metadata.New(map[string]string{})
	for k, v := range header {
		ret[strings.ToLower(k)] = v
	}
	return ret
}

func IsProtobufContentType(contentType string) bool {
	return strings.Contains(contentType, "protobuf")
}
