package impl

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

type OnFlushHandler func(data []byte) error

type LogBuffer struct {
	cacheSize      int
	onFlushHandler OnFlushHandler
	flushDuration  time.Duration
	data           []string
	start, count   int
	buf            *bytes.Buffer
	*sync.Mutex
}

func NewLogBuffer(cacheSize int, flushDuration time.Duration, handler OnFlushHandler) *LogBuffer {
	ret := &LogBuffer{
		cacheSize:      cacheSize,
		onFlushHandler: handler,
		flushDuration:  flushDuration,
		start:          0, count: 0,
		data:  make([]string, cacheSize),
		buf:   bytes.NewBuffer(nil),
		Mutex: &sync.Mutex{},
	}
	go ret.daemon()
	return ret
}

func (t *LogBuffer) daemon() {
	c := time.NewTicker(t.flushDuration)
	for range c.C {
		func() {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
				}
			}()
			if err := t.Flush(); err != nil {
				fmt.Println(err)
			}
		}()
	}
}

func (t *LogBuffer) reset() {
	t.start = 0
	t.count = 0
}

func (t *LogBuffer) checkFlush() error {
	if t.count >= t.cacheSize {
		if t.data != nil {
			if err := t.onFlushHandler(t.getData()); err != nil {
				return err
			}
		}
		t.reset()
	}
	return nil
}

func (t *LogBuffer) Append(data ...string) error {
	t.Lock()
	defer t.Unlock()
	if err := t.checkFlush(); err != nil {
		return err
	}
	for i := 0; i < len(data); i++ {
		next := (t.count + i) % t.cacheSize
		t.data[next] = data[i]
		t.count += 1
		if err := t.checkFlush(); err != nil {
			return err
		}
	}
	return nil
}

func (t *LogBuffer) getData() (ret []byte) {
	if len(t.data) == 0 {
		return make([]byte, 0)
	}
	t.buf.Reset()
	for i := t.start; i < t.start+t.count && i < t.cacheSize; i++ {
		t.buf.Write([]byte(t.data[i]))
	}
	d := t.cacheSize - t.start
	if d < t.count {
		for i := 0; i < t.count-d; i++ {
			t.buf.Write([]byte(t.data[i]))
		}
	}
	return t.buf.Bytes()
}

func (t *LogBuffer) Flush() error {
	t.Lock()
	defer t.Unlock()
	data := t.getData()
	defer t.reset()
	if len(t.data) > 0 && t.onFlushHandler != nil {
		return t.onFlushHandler(data)
	}
	return nil
}
