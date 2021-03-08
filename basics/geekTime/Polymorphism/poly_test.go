package poly_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type CppProgrammer struct {

}
func (c *CppProgrammer)WriteHelloWorld() Code {
	return "std::cout << \"Hello World!\" << std::endl;"
}

// interface 是一个指针
func clientWrite(p Programmer) {
	fmt.Printf("%T %v\n",p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goPro := &GoProgrammer{}
	cppPro := new(CppProgrammer)
	clientWrite(goPro)
	clientWrite(cppPro)
}
