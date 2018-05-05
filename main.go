package main

import (
	"fmt"
	"./services"
	"./httputils"
)

func main() {
	httputils.SetUpHTTPProxy("http://127.0.0.1:6152")
	fmt.Println("火币")
	fmt.Println(services.GetSymbols())
	fmt.Println(services.GetKLine("btcusdt", "15min", 1))
}
