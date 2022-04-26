package main

import "fmt"

func main(){
	var usd float32
	fmt.Scanln(&usd)
	
	var usdTobgn float32 = usd * 1.79549

	fmt.Println(usdTobgn)
}