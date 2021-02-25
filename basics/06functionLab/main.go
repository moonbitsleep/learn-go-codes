package main

import "fmt"

// 开氏温度转换到摄氏温度
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// 摄氏温度转换到华氏温度
func celsiusToFahrenheit(c float64) float64 {
	c = c*9.0/5.0 + 32.0
	return c
}

func kelvinToFahrenheit(k float64) float64 {
	k = kelvinToCelsius(k)
	k = celsiusToFahrenheit(k)
	return k
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%6.2f °K is %6.2f °C\n", kelvin, celsius)

	fah := kelvinToFahrenheit(kelvin)
	fmt.Printf("%6.2f °K is %6.2f °F\n", kelvin, fah)

	fah2 := celsiusToFahrenheit(celsius)
	fmt.Printf("%6.2f °C is %6.2f °F\n", celsius, fah2)

	k := 0.0
	f := kelvinToFahrenheit(k)
	fmt.Printf("%6.2f °K is %6.2f °F\n", k, f)
}
