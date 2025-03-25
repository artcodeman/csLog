package csLog

import (
	"github.com/artcodeman/csLog/file_log"
	"sync"
	"testing"
)

func TestLog_INFO(t *testing.T) {
	L := NewDefaultLog()
	defer func() {
		L.Close()
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

func TestFileLog_INFO(t *testing.T) {
	L := NewFileLog(file_log.SetFilePath("./"), file_log.SetStdin(true), file_log.SetDateFileType(file_log.DAY))
	defer func() {
		L.Close()

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
