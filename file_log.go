package csLog

import (
	defaultLog "github.com/artcodeman/csLog/default"
	fileLog "github.com/artcodeman/csLog/file_log"
	"github.com/artcodeman/csLog/log_base"
)

func NewFileLog(v ...fileLog.SetFileLogConfig) *Log {
	initServer := func() log_base.LogServer {
		v = append(v, fileLog.SetOpenFile(fileLog.OpenFile))
		s := fileLog.NewFileLogServer(v...)
		return s
	}
	initClient := func() log_base.LogClient {
		c := defaultLog.DefaultClient{}
		return &c
	}
	return NewLog(initClient, initServer)

}
