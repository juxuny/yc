package log

import (
	"fmt"
	"github.com/juxuny/yc/env"
	"github.com/juxuny/yc/trace"
	"runtime"
	"strings"
	"time"
)

const rpcLogTimeout = time.Millisecond * 10

type ColorFunc func(v interface{}) string

type RpcLogger struct {
	prefix       string
	level        Level
	ReportCaller bool
	depth        int
	appName      string
	logPrefix    string
	fields       map[string]interface{}
}

func NewRpcLogger() *RpcLogger {
	l := &RpcLogger{
		level:        currentLevel,
		appName:      env.GetString("LOG_SERVER_APP_NAME"),
		logPrefix:    env.GetString("LOG_SERVER_PREFIX"),
		ReportCaller: true,
	}
	return l
}

func NewRpcLoggerWithFields(fields map[string]interface{}) *RpcLogger {
	l := NewRpcLogger()
	l.fields = fields
	return l
}

func (l *RpcLogger) SetPrefix(prefix string) {
}

func (l *RpcLogger) SetCallStackDepth(depth int) *RpcLogger {
	l.depth = depth
	return l
}

func (l *RpcLogger) SetLevel(level Level) {
	l.level = level
}

func (l *RpcLogger) SetReportCaller(reportCaller bool) *RpcLogger {
	l.ReportCaller = reportCaller
	return l
}

func (l *RpcLogger) output(level string, v ...interface{}) string {
	var reqId = trace.GetReqId()
	var fieldList []string
	if uid := trace.GetUid(); uid.Valid {
		fieldList = append(fieldList, fmt.Sprintf("uid=%d", uid.Uint64))
	}
	if l.appName != "" {
		fieldList = append(fieldList, fmt.Sprintf("app=%s", l.appName))
	}
	if l.logPrefix != "" {
		fieldList = append(fieldList, fmt.Sprintf("prefix=%s", l.logPrefix))
	}
	if l.fields != nil {
		for k, v := range l.fields {
			fieldList = append(fieldList, fmt.Sprintf("%s=%v", k, v))
		}
	}
	var levelOutput = level
	if level == "ERROR" {
		levelOutput = Color.Red(level)
	}
	var messages = []string{
		fmt.Sprintf("[%s]", time.Now().Format("200601-02 15:04:05")),
		fmt.Sprintf("<reqId=%s>", reqId),
		fmt.Sprintf("[%s]", levelOutput),
	}
	if l.ReportCaller {
		_, file, line, ok := runtime.Caller(2 + l.depth)
		if ok {
			position := Color.LightPurple(fmt.Sprintf("(%s:%d)", file, line))
			messages = append(messages, position)
		}
	}
	messages = append(messages, strings.Join(fieldList, " "))
	var str = strings.Join(messages, " ") + " "
	var joinSlice = func(values ...interface{}) string {
		ret := ""
		for _, item := range values {
			ret += fmt.Sprintf("%v ", item)
		}
		return strings.Trim(ret, " ")
	}
	str += joinSlice(v...)
	return str
}

func (l *RpcLogger) Debug(v ...interface{}) {
	if l.level > LevelDebug {
		return
	}
	message := l.output("DEBUG", v...)
	fmt.Println(message)
}

func (l *RpcLogger) SQL(v ...interface{}) {
	if l.level > LevelInfo {
		return
	}
	message := l.output("INFO", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Consuming(v ...interface{}) {
	if l.level > LevelInfo {
		return
	}
	message := l.output("INFO", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Output(v ...interface{}) {
	message := l.output("INFO", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Debugf(format string, v ...interface{}) {
	if l.level > LevelDebug {
		return
	}
	message := l.output("DEBUG", fmt.Sprintf(format, v...))
	fmt.Println(message)
}

func (l *RpcLogger) Info(v ...interface{}) {
	message := l.output("INFO", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Infof(format string, v ...interface{}) {
	message := l.output("INFO", fmt.Sprintf(format, v...))
	fmt.Println(message)
}

func (l *RpcLogger) Print(v ...interface{}) {
	message := l.output("INFO", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Println(v ...interface{}) {
	message := l.output("INFO", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Printf(format string, v ...interface{}) {
	message := l.output("INFO", fmt.Sprintf(format, v...))
	fmt.Println(message)
}

func (l *RpcLogger) Warning(v ...interface{}) {
	message := l.output("WARN", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Warn(v ...interface{}) {
	message := l.output("WARN", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Warnf(format string, v ...interface{}) {
	message := l.output("WARN", fmt.Sprintf(format, v...))
	fmt.Println(message)
}

func (l *RpcLogger) Error(v ...interface{}) {
	message := l.output("ERROR", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Errorf(format string, v ...interface{}) {
	message := l.output("ERROR", fmt.Sprintf(format, v...))
	fmt.Println(message)
}

func (l *RpcLogger) Fatal(v ...interface{}) {
	message := l.output("ERROR", v...)
	fmt.Println(message)
}

func (l *RpcLogger) Fatalf(format string, v ...interface{}) {
	message := l.output("ERROR", fmt.Sprintf(format, v...))
	fmt.Println(message)
}

func (l *RpcLogger) Flush() {

}
