package empty_interface

import (
	"fmt"
	"testing"
)

func Dosomething(p interface{})  {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//}
	//if i,ok := p.(string); ok {
	//	fmt.Println("string", i)
	//}
	switch v := p.(type) {
	case int:
		fmt.Println("Integer ", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown Type")
	}
}


func TestEmptyInterfaceAssertion(t *testing.T) {
	Dosomething(10)
	Dosomething("233")
	Dosomething(100.8976)
}
