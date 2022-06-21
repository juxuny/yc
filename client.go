package yc

import (
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
	return t.entrypointCandidates[t.randInstance.Intn(len(t.entrypointCandidates))]
}

type RpcSignContentHandler interface {
	Sum(data []byte) (method SignMethod, signResult string, err error)
}
