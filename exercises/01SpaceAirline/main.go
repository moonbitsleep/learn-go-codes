package main

import (
	"fmt"
	"math/rand"
)

func drawSplitLine(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print("=")
	}
	fmt.Println("")
}

func drawLines(company string, day int, ttype string, price int) {
	fmt.Printf("%-17v%-6v%-11v%5v\n", company, day, ttype, price)
}

func drawAirlines() {
	fmt.Printf("%-17v%-6v%-11v%-5v\n", "Spaceline", "Days", "Trip type", "Price")
	drawSplitLine(39)
}

func genCompany() string {
	var company string
	var randCom = rand.Intn(3)
	switch randCom {
	case 0:
		company = "Virgin Galactic"
	case 1:
		company = "SpaceX"
	case 2:
		company = "Space Adventures"
	}
	return company
}

func randTType() string {
	var ttype string
	var randt = rand.Intn(2)
	switch randt {
	case 0:
		ttype = "One-single"
	case 1:
		ttype = "Round-trip"
	}
	return ttype
}

func genTicket() (string, int, string, int) {
	//2020年10月13日火星和地球的距离为 62,100,000 公里
	var distance = 62100000
	company := genCompany()
	var resultDay int
	var resultPrice int
	// 生成每秒随机速度
	var speedSecond = rand.Intn(30-16) + 1 + 16
	//换算成每天的速度
	var speedDay = speedSecond * 3600 * 24
	// 求得单程所需天数
	var singleTripDay = distance / speedDay
	// 求得每秒1公里的速度需要多少百万美元
	var priceSpeed = (50 - 36) / (30 - 16)
	// 求得单程价格
	var priceSingle = priceSpeed * speedSecond

	ttype := randTType()
	if ttype == "One-single" {
		resultDay = singleTripDay
		resultPrice = priceSingle
	} else {
		resultDay = singleTripDay * 2
		resultPrice = priceSingle * 2
	}

	return company, resultDay, ttype, resultPrice
}

func drawTickets() {
	drawAirlines()
	for i := 0; i < 10; i++ {
		company, day, ttype, price := genTicket()
		drawLines(company, day, ttype, price)
	}
}
func main() {
	drawTickets()
}
