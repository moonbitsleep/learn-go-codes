package func_test

import (
	"fmt"
	"testing"
	"time"
)

func timeSpend(inner func (op int) int) func (op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpend(slowFunc)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	sum := 0
	for _, v := range ops {
		sum += v
	}
	return sum
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear() {
	fmt.Println("Clear resources.")
}


func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start!")
	fmt.Println("Processing...")
	fmt.Println("Closeing...")
}
