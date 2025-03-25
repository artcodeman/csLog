package log_base

type LogType interface {
	GetLog() string
}

type LogClient interface {
	INFO(...any) string
	WARN(...any) string
	DEBUG(...any) string
	ERROR(...any) string
}

type LogServer interface {
	Out(log string)
	Close()
	Print(log string)
	GetLog() (string, bool)
}

type LogServerInitFunc func() LogServer

type LogClientInitFunc func() LogClient
