package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type cat struct {
	name string
}

func (c cat) talk() string {
	return "mow mow mow"
}

type duoLa int

func (a duoLa) talk() string{
	return strings.Repeat("DaXiong ", int(a))
}

type comCat struct {
	cat
}

func main() {
	var t talker
	t = cat{name: "white"}
	fmt.Println(t.talk())

	t = duoLa(4)
	fmt.Println(t.talk())

	t = comCat{}
	shout(t)


}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}