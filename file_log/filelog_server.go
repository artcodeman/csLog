package file_log

import (
	"fmt"
	defaultLog "github.com/artcodeman/csLog/default"
	"os"
	"path"
	"strconv"
	"time"
)

type FileLogServer struct {
	defaultLog.DefaultServer
	f            *os.File
	NowFptr      *time.Time
	path         string
	DateFileType DateFileType
	stdin        bool
	Step         int              //频次
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
	d.Lock()
	d.openFile(d)
	if d.stdin {
		fmt.Println(log)
	}
	if d.f != nil {
		_, err := d.f.Write([]byte(log + "\n"))
		if err != nil {
			println("write err %v", err)
		}
	}
	d.Unlock()

}

type SetFileLogConfig func(d *FileLogServer)

// SetFile 设置输出文件
func SetFile(f *os.File) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.f = f
	}
}

// SetFilePath 设置输出Pathz
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

// SetOpenFile 设置打开文件的函数
func SetOpenFile(f SetFileLogConfig) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.openFile = f
	}
}

func SetDateFileType(ft DateFileType) SetFileLogConfig {
	return func(d *FileLogServer) {
		d.DateFileType = ft
	}
}

func NewFileLogServer(v ...SetFileLogConfig) *FileLogServer {
	var s = new(FileLogServer)
	s.Init()
	for _, f := range v {
		f(s)
	}
	return s
}

func OpenFile(d *FileLogServer) {
	if d.NowFptr == nil {
		t := time.Now()
		d.NowFptr = &t
	}
	fileName := fmt.Sprintf("%d.log", d.NowFptr.Unix())
	filepathStr := path.Join(d.path, fmt.Sprintf("%d.log", d.NowFptr.Unix()))
	if d.f != nil {
		fileUnix, _ := strconv.ParseInt(d.f.Name()[:10], 10, 64)
		switch d.DateFileType {
		case NONE:
			break
		case HOUR:
			if d.NowFptr.Unix()-fileUnix > int64(time.Hour.Seconds()) {
				d.f.Close()
				d.f = nil
			}
		case DAY:
			if d.NowFptr.Unix()-fileUnix > int64(time.Hour.Seconds()*24) {
				d.f.Close()
				d.f = nil
			}
		case WEEK:
			if d.NowFptr.Unix()-fileUnix > int64(time.Hour.Seconds()*24*7) {
				d.f.Close()
				d.f = nil
			}
		case MONTH:
			if d.NowFptr.Unix()-fileUnix > int64(time.Hour.Seconds()*24*30) {
				d.f.Close()
				d.f = nil
			}
		default:
			break
		}
	}
	if d.f == nil {
		if _, err := os.Stat(filepathStr); err != nil {
			os.Chdir(d.path)
			d.f, _ = os.Create(fileName)
		} else {
			d.f, _ = os.Open(fileName)
		}
	}

}
