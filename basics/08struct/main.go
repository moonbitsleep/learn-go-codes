package main

import "fmt"
import "math"

type location struct {
	x float64
	y float64
}

func distance(loc1, loc2 location) float64 {
	ret := location{0.0, 0.0}
	ret.x = loc2.x - loc1.x
	ret.y = loc2.y - loc1.y
	dis := math.Sqrt(math.Pow(ret.x, 2.0) + math.Pow(ret.y, 2.0))
	return dis
}

func main() {
	p1 := location{x: 1, y: 2}
	p2 := location{x: 4, y: 8}
	p12Dis := distance(p1, p2)
	fmt.Println(p12Dis)
}



