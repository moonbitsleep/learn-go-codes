package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	defer func() {
		// 谨慎使用，避免僵尸程序
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something went wrong!"))
}
