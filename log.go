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

func start(s log_base.LogServer) {
	go func() {
		defer func() {
			println("已退出")
		}()
		for {
			v, ok := s.GetLog()
			if v == "" && !ok {
				return
			}
			if ok {
				s.Print(v)
			}
		}
	}()

}

func NewLog(logClientFunc log_base.LogClientInitFunc, logServerFunc log_base.LogServerInitFunc) *Log {
	logServer := logServerFunc()
	start(logServer)
	return &Log{logClient: logClientFunc(), logServer: logServer}
}
