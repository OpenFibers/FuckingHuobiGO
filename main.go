package main

import (
	"fmt"
	"./services"
)

func main() {
	fmt.Println("火币")
	fmt.Println(services.GetSymbols())
	fmt.Println(services.GetKLine("btcusdt", "15min", 1))
}
