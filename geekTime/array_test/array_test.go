package array_test

import (
	"fmt"
	"testing"
)

func TestArraySelection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arr_sec := arr[1:]
	t.Log(arr_sec)
	fmt.Printf("%T--%v\n", arr_sec, arr_sec)
}
