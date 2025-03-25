package defaultLog

import (
	"sync"
	"sync/atomic"
	"time"
)

type DefaultServer struct {
	p       chan string
	isClose atomic.Int32
	sync.Mutex
}

func (d *DefaultServer) GetLog() (string, bool) {
	v, ok := <-d.p
	if !ok {
		d.isClose.Add(1)
	}
	return v, ok
}

func (d *DefaultServer) Print(log string) {
	println(log)
}

func (d *DefaultServer) Out(msg string) {
	if d.isClose.Load() == 0 {
		d.p <- msg
	}
}
func (d *DefaultServer) Close() {
	if d.isClose.Load() == 0 {
		close(d.p)
		d.isClose.Add(1)
	}
	for d.isClose.Load() == 1 {
		time.Sleep(100 * time.Microsecond)
	}

}

func (d *DefaultServer) Init() {
	d.p = make(chan string, 10)

}
