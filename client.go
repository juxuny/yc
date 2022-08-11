package yc

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

type EntrypointDispatcher interface {
	SelectOne() string
}

type RandomEntrypointDispatcher struct {
	randInstance         *rand.Rand
	entrypointCandidates []string
}

func NewRandomEntrypointDispatcher(entrypointCandidates []string) EntrypointDispatcher {
	return &RandomEntrypointDispatcher{
		randInstance:         rand.New(rand.NewSource(time.Now().UnixNano())),
		entrypointCandidates: entrypointCandidates,
	}
}

func (t *RandomEntrypointDispatcher) SelectOne() string {
	if len(t.entrypointCandidates) == 0 {
		return ""
	}
	t.randInstance.Seed(time.Now().UnixNano())
	return t.entrypointCandidates[t.randInstance.Intn(len(t.entrypointCandidates))]
}

type RpcSignContentHandler interface {
	Sum(data []byte) (method SignMethod, signResult string, err error)
	GetAccessKey() string
}

type Sha256SignHandler struct {
	accessKey string
	secret    string
}

func (t *Sha256SignHandler) Sum(data []byte) (method SignMethod, signResult string, err error) {
	h := sha256.New()
	return SignMethodSha256, fmt.Sprintf("%02x", h.Sum(data)), nil
}

func (t *Sha256SignHandler) GetAccessKey() string {
	return t.accessKey
}

func NewDefaultSignHandler(accessKey string, secret string) RpcSignContentHandler {
	return &Sha256SignHandler{
		accessKey: accessKey,
		secret:    secret,
	}
}
