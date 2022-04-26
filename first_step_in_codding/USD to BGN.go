// Write a program for converting US dollars (USD) into Bulgarian levs (BGN). You will receive a real number from the console, which will show the number of US dollars. 
// Use a fixed exchange rate when converting between dollar and lev: 1 USD = 1.79549 BGN.

package main

import "fmt"

func main(){
	var usd float32
	fmt.Scanln(&usd)
	
	var usdTobgn float32 = usd * 1.79549

	fmt.Println(usdTobgn)
}
