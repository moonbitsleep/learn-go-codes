package main

import "fmt"

func traversMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	m := map[string]string {
		"name" : "ravi",
		"sex" : "male",
		"age":"22",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	traversMap(m)

	fmt.Println("Getting values...")
	sex, ok := m["sex"]
	fmt.Println(sex, ok)

	if addr, ok := m["addr"]; ok {
		fmt.Println(addr)
	} else {
		fmt.Println("key does not exist...")
	}

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("%q, %v",name, ok)
}
