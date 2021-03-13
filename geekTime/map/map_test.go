package map_test

import (
	"testing"
)

func TestInMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "four": 4}
	t.Logf("len m1 = %d\n", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 = %d\n", len(m2))

	// 初始化的是cap而不是len
	m3 := make(map[string]int, 10)
	t.Logf("len m3 = %d\n", len(m3))
	// cap函数不能求map
}
