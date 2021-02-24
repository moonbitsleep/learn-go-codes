package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	const secondPerDay = 86400

	const days = distance / lightSpeed / secondPerDay

	fmt.Println("Canis Major Dwarf is ", days, "light days away")
}
