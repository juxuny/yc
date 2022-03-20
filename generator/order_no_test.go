package generator

import (
	"strings"
	"testing"
	"time"
)

func TestNewOrderNoGenerator(t *testing.T) {
	generator := NewOrderNoGenerator(1)
	l := 100000
	ret := make([]string, l)
	m := make(map[string]bool)
	start := time.Now()
	for i := 0; i < l; i++ {
		ret[i] = generator.Gen()
		m[ret[i]] = true
	}
	end := time.Now()
	t.Log(strings.Join(ret, ","))
	t.Log(end.Sub(start))
	t.Log(len(m))
	if len(m) != l {
		t.Fatal("result duplicated")
	}
}
