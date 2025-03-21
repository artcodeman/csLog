package csLog

import "github.com/artcodeman/csLog/log_base"

type Log struct {
	logClient log_base.LogClient
	logServer log_base.LogServer
}

func (receiver *Log) INFO(M ...any) {
	log := receiver.logClient.INFO(M...)
	receiver.logServer.Out(log)

}

func (receiver *Log) ERROR(M ...any) {
	log := receiver.logClient.ERROR(M...)
	receiver.logServer.Out(log)
}

func (receiver *Log) WARN(M ...any) {
	log := receiver.logClient.WARN(M...)
	receiver.logServer.Out(log)

}

func (receiver *Log) DEBUG(M ...any) {
	log := receiver.logClient.DEBUG(M...)
	receiver.logServer.Out(log)
}

func NewLog(logClientFunc log_base.LogClientInitFunc, logServerFunc log_base.LogServerInitFunc) *Log {
	return &Log{logClient: logClientFunc(), logServer: logServerFunc()}
}
