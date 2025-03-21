package csLog

type DefaultServer struct {
	path string
	p    chan string
}

func (d *DefaultServer) Print(log string) {
	println(log)
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

func (d *DefaultServer) Out(msg string) {
	d.p <- msg

}

func (d *DefaultServer) FilePath(path string) {
	d.path = path
}

func (d *DefaultServer) Close() {
	close(d.p)
}
