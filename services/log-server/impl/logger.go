package impl

type Logger interface {
	Info(app string, msg string) error
	Flush() error
}
