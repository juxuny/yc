package log

import (
	"fmt"
	"log"
)

type Level int

const (
	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
	LevelFatal = 5
	LevelNone  = 10000 // 不输出
)

var (
	currentLevel = Level(LevelDebug)
)

func SetLevel(lv Level) {
	currentLevel = lv
}

var (
	enableSql = true
)

type ILogger interface {
	SetPrefix(string)
	Println(...interface{})
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Debug(...interface{})
	Printf(string, ...interface{})
	Print(...interface{})
	SQL(...interface{})
	Consuming(...interface{})
	Output(...interface{})
	Flush()
}

var DefaultLogger = NewLogger("[DEFAULT] ")

type LoggerV1 struct {
	l         *log.Logger
	CallDepth int
}

func (t *LoggerV1) Flush() {
}

func (t *LoggerV1) SetPrefix(s string) {
	t.l.SetPrefix(s)
}

func (t *LoggerV1) Println(v ...interface{}) {
	t.Output("[INFO] " + fmt.Sprint(v...))
}

func (t *LoggerV1) Error(v ...interface{}) {
	if currentLevel <= LevelError {
		t.Output("[ERROR] \033[0;31m" + fmt.Sprint(v...) + "\033[0m")
	}
}

func (t *LoggerV1) Warn(v ...interface{}) {
	if currentLevel <= LevelWarn {
		t.Output("[WARN] \033[0;33m" + fmt.Sprint(v...) + "\033[0m")
	}
}

func (t *LoggerV1) Info(v ...interface{}) {
	if currentLevel <= LevelInfo {
		t.Output("[INFO] \033[0;32m" + fmt.Sprint(v...) + "\033[0m")
	}
}

func (t *LoggerV1) Debug(v ...interface{}) {
	if currentLevel <= LevelDebug {
		t.Output("[DEBUG] " + fmt.Sprint(v...) + "")
	}
}

func (t *LoggerV1) Printf(format string, v ...interface{}) {
	if currentLevel <= LevelInfo {
		t.Output("[INFO] " + fmt.Sprintf(format, v...))
	}
}

func (t *LoggerV1) SQL(v ...interface{}) {
	if enableSql {
		t.Output("[SQL] \033[0;35m" + fmt.Sprint(v...) + "\033[0m")
	}
}

func (t *LoggerV1) Consuming(v ...interface{}) {
	if enableSql {
		t.Output("[CONSUMING] \033[0;36m" + fmt.Sprint(v...) + "\033[0m")
	}
}

//func NewLogger(prefix string, callDepth ...int) ILogger {
//	cd := 3
//	if len(callDepth) > 0 {
//		cd = callDepth[0]
//	}
//	ret := &LoggerV1{
//		l:         log.New(os.Stdout, prefix+" ", log.Ltime|log.Llongfile|log.Ldate|log.LstdFlags),
//		CallDepth: cd,
//	}
//	return ret
//}

func NewLogger(prefix string, callDepth ...int) ILogger {
	cd := 0
	if len(callDepth) > 0 {
		cd = callDepth[0]
	}
	ret := NewRpcLogger()
	ret.SetPrefix(prefix)
	ret.SetReportCaller(true)
	ret.SetCallStackDepth(cd)
	return ret
}

func (t *LoggerV1) Print(v ...interface{}) {
	t.Output("[INFO] " + fmt.Sprint(v...))
}

func (t *LoggerV1) Output(v ...interface{}) {
	_ = t.l.Output(t.CallDepth, fmt.Sprint(v...))
}
