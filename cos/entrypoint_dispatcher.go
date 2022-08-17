package cos

import (
	"fmt"
	"github.com/juxuny/yc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type entrypointDispatcher struct {
	*client
	*sync.RWMutex
	entrypointCandidates []string
	randInstance         *rand.Rand
}

func (t *entrypointDispatcher) SelectOne() string {
	if len(t.entrypointCandidates) == 0 {
		return ""
	}
	return t.entrypointCandidates[t.randInstance.Intn(len(t.entrypointCandidates))]
}

func NewEntrypointDispatcher(key string, opt Options) yc.EntrypointDispatcher {
	client := NewClient(opt)
	configMap, err := client.fetch()
	if err != nil {
		panic(err)
	}
	values := make([]string, 0)
	if vs, b := configMap[key]; b {
		values = strings.Split(vs, ",")
	} else {
		panic(fmt.Sprintf("not found key '%s' in cos configId: %s", key, opt.ConfigId))
	}
	return &entrypointDispatcher{
		client:               NewClient(opt),
		RWMutex:              &sync.RWMutex{},
		entrypointCandidates: values,
		randInstance:         rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
