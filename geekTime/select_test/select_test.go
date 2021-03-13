package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	//retCh := make(chan string)
	// buffer channel 写法 放数据未完成便可以退出这个协程
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result.")
		// 向channel 放数据
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
		case ret := <-AsyncService():
			t.Log(ret)
		case <- time.After(time.Millisecond * 100):
			t.Error("Time out")
	}
}