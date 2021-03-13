package condition_test

import "testing"

func TestSwitchCaseCondition(t *testing.T) {
	//switch 代替 ifel
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unknow")
		}
	}
}
