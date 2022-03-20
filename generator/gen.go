package generator

import (
	"strings"
	"time"
)

// ms
var startEndPoint int64

func init() {
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-03-10 00:00:00", time.Local)
	startEndPoint = startTime.UnixNano() / int64(time.Millisecond)
}

var globalNodeId int

var orderNoGenerator *OrderNoGenerator

func Init(nodeId int) {
	globalNodeId = nodeId
	orderNoGenerator = NewOrderNoGenerator(nodeId)
}

func GenerateOrderNo(prefix ...string) string {
	if orderNoGenerator == nil {
		panic("please call Init(nodeId) first")
	}
	orderNo := orderNoGenerator.Gen()
	if len(prefix) > 0 {
		return strings.Join(prefix, "_") + orderNo
	}
	return orderNo
}
