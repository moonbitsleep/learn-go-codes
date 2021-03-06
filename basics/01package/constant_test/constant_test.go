package constant_test

import "testing"

const (
	Mon = iota + 1
	Tus
	Wed
	Thu
	Fri
	Sat
	Sun
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Mon, Tus, Wed, Thu, Fri, Sat, Sun)
}

func TestConstant1(t *testing.T) {

	t.Log(Readable, Writable, Executable)
}
