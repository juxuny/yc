package cos

import (
	"crypto/sha256"
	"fmt"
	"github.com/juxuny/yc"
)

type sha256SignHandler struct {
	accessKey string
	secret    string
}

func (t *sha256SignHandler) Sum(data []byte) (method yc.SignMethod, signResult string, err error) {
	h := sha256.New()
	return yc.SignMethodSha256, fmt.Sprintf("%02x", h.Sum(data)), nil
}

func (t *sha256SignHandler) GetAccessKey() string {
	return t.accessKey
}

func NewDefaultSignHandler(opt Options) yc.RpcSignContentHandler {
	type CosSignConfig struct {
		AccessKey string `cos:"client_access_key"`
		Secret    string `cos:"client_secret"`
	}
	client := NewClient(opt)
	configMap, err := client.fetch()
	if err != nil {
		panic(err)
	}
	var cosSignConfig CosSignConfig
	if err := Parse(configMap, &cosSignConfig); err != nil {
		panic(err)
	}
	return &sha256SignHandler{
		accessKey: cosSignConfig.AccessKey,
		secret:    cosSignConfig.Secret,
	}
}
