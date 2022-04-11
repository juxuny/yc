package router

import (
	"encoding/json"
	"net/http"
)

func WriteString(w http.ResponseWriter, data string, code ...int) {
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(data))
}

func WriteJson(w http.ResponseWriter, data interface{}, code ...int) {
	w.Header().Set("Content-Type", "application/json;utf8")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	jsonData, _ := json.Marshal(data)
	_, _ = w.Write(jsonData)
}

func WriteSuccessData(w http.ResponseWriter, data interface{}) {
	WriteJson(w, map[string]interface{}{
		"code": 0,
		"data": data,
	})
}
