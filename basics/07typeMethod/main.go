package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

func (k kelvin) Celsius() celsius {
	return celsius(k - 273.15)
}

func (c celsius) Fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}

func (k kelvin) Fahrenheit() fahrenheit {
	var c celsius = k.Celsius()
	var f fahrenheit = c.Fahrenheit()
	return f
}

func (f fahrenheit) Celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k kelvin = 233.0
	var f fahrenheit = k.Fahrenheit()
	fmt.Printf("%6.2f °K is %6.2f °F\n", k, f)

	var c celsius = 37.0
	f = c.Fahrenheit()
	fmt.Printf("%6.2f °C is %6.2f °F\n", c, f)

	f = -40.27
	c = f.Celsius()
	fmt.Printf("%6.2f °C is %6.2f °F\n", c, f)

	c = k.Celsius()
	fmt.Printf("%6.2f °C is %6.2f °K\n", c, k)
}
