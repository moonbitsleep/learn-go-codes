package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	// map a function
	m := map[int]func(op int) int {}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op * op}
	m[3] = func(op int) int {return op * op * op}
	t.Log(m[1](2), m[2](2), m[3](2))

}

// Set unique elem

func TestMapForSet(t *testing.T) {
	var mySet = map[int]bool {}
	var a = []int{1, 2, 2, 3, 5, 6, 8, 5, 3, 8, 9}

	for _,v :=range a {
		if _, ok := mySet[v]; !ok {
			mySet[v] = true
		}
	}
	t.Log(mySet)
	delete(mySet, 90)
	t.Log(mySet)
}