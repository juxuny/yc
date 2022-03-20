package generator

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type OrderNoGenerator struct {
	*sync.Mutex
	index         int64
	lastTimestamp int64
	nodeId        int
}

func NewOrderNoGenerator(nodeId int) *OrderNoGenerator {
	return &OrderNoGenerator{
		Mutex:         &sync.Mutex{},
		index:         0,
		lastTimestamp: time.Now().UnixNano() / int64(time.Millisecond),
		nodeId:        nodeId,
	}
}

func (t *OrderNoGenerator) Gen() string {
	var current time.Time
	t.Lock()
	for {
		current = time.Now()
		timestamp := current.UnixNano() / int64(time.Millisecond)
		if timestamp == t.lastTimestamp && (t.index+1)%1000 != 0 {
			t.index += 1
			t.index %= 1000
			t.lastTimestamp = timestamp
			break
		} else if timestamp != t.lastTimestamp {
			t.index = 0
			t.lastTimestamp = timestamp
		}
	}
	t.Unlock()
	ret := current.Format("20060102150405.000")
	ret = strings.ReplaceAll(ret, ".", "")
	ret = ret + fmt.Sprintf("%03d%03d", t.nodeId, t.index)
	return ret
}
