package err_t_test

import (
	"errors"
	"testing"
)

// error 接口仅包含一个返回字符型的方法
// 快速失败
func getFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2, 100]")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i - 2] + fibList[i - 1])
	}
	return fibList, nil
}


func TestGetFib(t *testing.T) {
	if v, err := getFibonacci(76); err != nil {
		t.Error(err)
	}else {
		t.Log(v)
	}

}
