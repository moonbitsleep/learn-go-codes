package async_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

// 串行操作
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
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

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	// 向channel 取数据
	fmt.Println(<-retCh)
	// 被阻塞
	time.Sleep(time.Second * 1)
}