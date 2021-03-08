package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 1; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Millisecond * 5)
}
