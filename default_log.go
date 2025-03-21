package csLog

import (
	defaultLog "github.com/artcodeman/csLog/default"
	"github.com/artcodeman/csLog/log_base"
)

func NewDefaultLog() *Log {
	initServer := func() log_base.LogServer {
		s := defaultLog.DefaultServer{}
		s.Init()
		return &s
	}
	initClient := func() log_base.LogClient {
		c := defaultLog.DefaultClient{}
		return &c
	}
	return NewLog(initClient, initServer)

}
