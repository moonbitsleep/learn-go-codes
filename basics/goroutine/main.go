package main

import (
	"fmt"
	"time"
)

func sleepGopher()  {
	fmt.Println("呼呼呼……")
	time.Sleep(time.Second * 2)
	fmt.Println("睡醒了！")
}

func main()  {
	/***
	// 开辟协程并行运行此函数
	go sleepGopher()
	//主函数结束前没执行玩的必须结束
	//time.Sleep(time.Second * 1) // 呼呼呼…… 没有醒
	time.Sleep(time.Second * 3) // 呼呼呼…… 睡醒了！
	 */
	for i := 1; i < 5; i++ {
		go wakeGopher(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("等待5s，执行结束!")
}

func wakeGopher(i int) {
	fmt.Printf("呼呼呼……%d号gopher在打鼾\n", i)
	time.Sleep(time.Second * 1)
	fmt.Printf("%d号gopher睡醒了！\n", i)
}

/**
呼呼呼……4号gopher在打鼾
呼呼呼……1号gopher在打鼾
呼呼呼……3号gopher在打鼾
呼呼呼……2号gopher在打鼾
2号gopher睡醒了！
1号gopher睡醒了！
3号gopher睡醒了！
4号gopher睡醒了！
等待5s，执行结束!

这种方法很不方便！浪费时间！
 */