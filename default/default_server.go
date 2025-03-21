package defaultLog

type DefaultServer struct {
	p chan string
}

func (d *DefaultServer) Print(log string) {
	println(log)
}

func (d *DefaultServer) Out(msg string) {
	d.p <- msg

}
func (d *DefaultServer) Close() {
	close(d.p)
}

func (d *DefaultServer) start() {
	go func() {
		defer func() {
			println("已退出")
		}()
		for {
			v, ok := <-d.p
			if v == "" && !ok {
				return
			}
			if ok {
				d.Print(v)
			}
		}
	}()

}

func (d *DefaultServer) Init() {
	d.p = make(chan string, 10)
	d.start()

}
