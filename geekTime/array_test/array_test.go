package array_test

import (
	"fmt"
	"testing"
)

func TestArraySelection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arrSec := arr[1:]
	t.Log(arrSec)
	fmt.Printf("%T--%v\n", arrSec, arrSec)
}
