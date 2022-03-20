package log

type gormLogger struct {
	*RpcLogger
}

func NewGormLogger() *gormLogger {
	l := gormLogger{
		RpcLogger: NewRpcLogger().SetReportCaller(false),
	}
	return &l
}

// Print format & print log
func (l gormLogger) Print(values ...interface{}) {
	l.RpcLogger.Print(gormLogFormatter(values...)...)
}
