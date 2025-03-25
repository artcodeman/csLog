# CsLog


[![license](https://img.shields.io/github/license/:user/:repo.svg)](LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/artcodeman/csLog)
中文 ｜  [English ](README_EN.md)

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
```go
L := NewFileLog(
	file_log.SetFilePath("./"), 
	file_log.SetStdin(true), 
	file_log.SetDateFileType(file_log.DAY)
	)
L.INFO("sssssssss")
```

setFilePath: 输出目录\
setStdin:同时输出到标志输出流\
SetDateFileType: 日志文件分割周期,支持按月,周,日,小时,默认不分割

### Tips

由于日志打印服务是异步的,当程序退出时应调用close方法关闭日志输出服务,否则可能导致日志丢失！！！
```
//Example:
defer func() { L.Close() }()
```


### 高级用法

可以嵌套DefaultServer,重写Print方法,将日志内容输出到指定位置，或发送到指定服务进行存储\
可参考FileLogServer的实现
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

待补充



## License

[MIT © Richard McRichface.](./LICENSE)