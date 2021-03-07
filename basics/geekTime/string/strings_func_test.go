package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "How Are You"
	t.Logf("s is %[1]T type --- %[1]v\n",s)
	parts := strings.Split(s, " ")
	t.Logf("parts is %[1]T type --- %[1]v\n",parts)
	for _, part := range parts {
		t.Log(part)
	}

	whole := strings.Join(parts, "?")
	t.Log(whole)
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(1010101100)
	t.Log("str: " + s)
	if i, err := strconv.Atoi(s); err == nil {
		t.Log(i + 90)
	}
}
