package operator_test

import "testing"

func TestCompareArry(t *testing.T) {
	a := [...]int{1, 2, 4, 5}
	c := [...]int{1, 2, 4, 5}
	b := [...]int{2, 8, 10, 12}
	t.Log("a == c", a == c)
	t.Log("a == b", a == b)
}
