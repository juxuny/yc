package log_server

func (t LogLevel) Name() string {
	if t == LogLevel_Debug {
		return "DEBUG"
	} else if t == LogLevel_Info {
		return "INFO"
	} else if t == LogLevel_Warn {
		return "WARN"
	} else if t == LogLevel_Error {
		return "ERROR"
	} else {
		return "UNKNOWN"
	}
}
