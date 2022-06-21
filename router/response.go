package router

import (
	"encoding/json"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/log"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func WriteString(w http.ResponseWriter, data string, code ...int) {
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set(HeaderContentType, "text/plain")
	_, _ = w.Write([]byte(data))
}

func WriteJson(w http.ResponseWriter, data interface{}, code ...int) {
	w.Header().Set(HeaderContentType, "application/json;utf8")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	var jsonData []byte
	jsonData, _ = json.Marshal(data)
	_, _ = w.Write(jsonData)
}

func WriteSuccessData(w http.ResponseWriter, data interface{}) {
	WriteJson(w, map[string]interface{}{
		"code": 0,
		"data": data,
	})
}

func WriteProtobuf(w http.ResponseWriter, data proto.Message, code ...int) {
	w.Header().Set(HeaderContentType, "application/protobuf")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	respBody, err := proto.Marshal(data)
	if err != nil {
		log.Error(err)
	}
	_, err = w.Write(respBody)
	if err != nil {
		log.Error(err)
	}
}

func WriteProtobufError(w http.ResponseWriter, errContent *dt.Error, code ...int) {
	w.Header().Set(HeaderContentType, "application/protobuf")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	respBody, err := proto.Marshal(errContent)
	if err != nil {
		log.Error(err)
	}
	_, err = w.Write(respBody)
	if err != nil {
		log.Error(err)
	}
}
