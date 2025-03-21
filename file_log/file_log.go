package file_log

import (
	defaultLog "github.com/artcodeman/csLog/default"
	"os"
	"time"
)

type FileLogServer struct {
	defaultLog.DefaultServer
	f            *os.File
	nowFptr      *time.Time
	path         string
	DateFileType DateFileType
	stdin        bool
	step         int              //频次
	openFile     SetFileLogConfig //openFile方法
}

type DateFileType int

const (
	NONE DateFileType = iota
	HOUR
	DAY
	WEEK
	MONTH
)

func (d *FileLogServer) Print(log string) {
	d.openFile(d)
	if d.stdin {
		println(log)
	}
	if d.f != nil {
		_, err := d.f.Write([]byte(log))
		if err != nil {
			println("write err %v", err)
			return
		}
	}

}

type SetFileLogConfig func(d *FileLogServer)

// SetFile 设置输出文件
func SetFile(f *os.File) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.f = f
	}
}

// SetFilePath 设置输出Path
func SetFilePath(path string) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.path = path
	}
}

// SetStdin 是否输出到控制胎
func SetStdin(stdin bool) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.stdin = stdin
	}
}

func SetOpenFile(f SetFileLogConfig) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.openFile = f
	}
}
