package main

import "fmt"

type report struct {
	sol int // 日期
	// 类型嵌入
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius{
	return (t.high + t.low) / 2
}

func main() {
	report := report{
		sol: 15,
		location: location{-4.1468, 55.98728},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	// 转发方法
	fmt.Println(report.temperature.average(), report.average())
	//外部结构访问内部变量
	fmt.Println(report.low)
	report.low = -98.0
	fmt.Println(report.temperature.low)
}
