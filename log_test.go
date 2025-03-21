package csLog

import (
	"sync"
	"testing"
)

func TestLog_INFO(t *testing.T) {
	L := NewDefaultLog()
	defer func() {
		L.Server.Close()

	}()
	wait := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wait.Add(1)
		i := i
		go func() {
			L.INFO(i, "info sssssssss")
			wait.Done()
		}()
	}
	wait.Wait()
	return

}
