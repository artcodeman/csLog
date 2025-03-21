package csLog

import "github.com/artcodeman/csLog/log_base"

type Log struct {
	LogClient log_base.LogClient
	Server    log_base.LogServer
}

func (receiver *Log) INFO(M ...any) {
	log := receiver.LogClient.INFO(M...)
	receiver.Server.Out(log)

}

func (receiver *Log) ERROR(M ...any) {
	log := receiver.LogClient.ERROR(M...)
	receiver.Server.Out(log)
}

func (receiver *Log) WARN(M ...any) {
	log := receiver.LogClient.WARN(M...)
	receiver.Server.Out(log)

}

func (receiver *Log) DEBUG(M ...any) {
	log := receiver.LogClient.DEBUG(M...)
	receiver.Server.Out(log)
}

func NewLog(logClientFunc log_base.LogClientInitFunc, logServerFunc log_base.LogServerInitFunc) *Log {
	return &Log{LogClient: logClientFunc(), Server: logServerFunc()}
}
