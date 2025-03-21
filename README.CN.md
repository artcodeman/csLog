# CsLog


[![license](https://img.shields.io/github/license/:user/:repo.svg)](LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/artcodeman/csLog)
中文 ｜  [English ](README.md)

csLog是一个线程安全的异步日志库，提供四种日志级别,可高度自定义化

## 目录

- [安装](#安装)
- [我该如何使用](#如何使用)
- [相关API](#api)

## 安装


```
get go get github.com/artcodeman/csLog
```


## 如何使用
 csLog提供了两种日志打印服务，DefaultServer,FileLogServer\
 DefaultServer : 提供简单的日志打印队列服务\
 FileLogServer : 提供日志输入文件相关服务
### 基础用法

### 高级用法

可以嵌套DefaultServer,重写Print方法,将日志内容输出到指定位置，或发送到指定服务进行存储
```
//Example:
type FileLogServer struct {
	defaultLog.DefaultServer
	f        *os.File
	nowFptr  *time.Time
	path     string
	fileType int
	stdin    bool
	step     int
	openFile func()
}

type DateFileType int

const (
	HOUR DateFileType = iota
	DAY
	WEEK
	MONTH
)

func (d *FileLogServer) Print(log string) {
	d.openFile()
	if d.stdin {
		println(log)
	}
	if d.f!=nil{
		_, err := d.f.Write([]byte(log))
		if err != nil {
			println("write err %v",err)
			return 
		}
	}

}


```



## API





## License

[MIT © Richard McRichface.](./LICENSE)