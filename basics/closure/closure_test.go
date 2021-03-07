package closure

import (
	"fmt"
	"strings"
	"testing"
)

func makeSuffix(suffix string) func (name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}

}

func normalMakeSufffix(suffix string, name string) string {
	if strings.HasSuffix(name, suffix) == false {
		return name + suffix
	}
	return name
}

func TestClosure(t *testing.T) {
	f := makeSuffix(".jpg")

	fmt.Println(f("winter"))
	fmt.Println(f("win.jpg"))
	fmt.Println(normalMakeSufffix(".txt", "actor"))
	fmt.Println(normalMakeSufffix(".txt", "actress.avi"))
}
