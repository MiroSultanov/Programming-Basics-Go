// Lily is already N years old. She receives a gift for each of her birthdays.
// • For odd birthdays (1, 3, 5 ... n) receives toys.
// • For even birthdays (2, 4, 6 ... n) receives money.
// For the second birthday he receives BGN 10.00, as the amount increases by BGN 10.00, for each subsequent even birthday (2 -> 10, 4 -> 20, 6 -> 30 ... etc.). 
// Over the years, Lily has secretly saved money. Lily's brother, in the years she receives money, takes BGN 1.00 from them. Lily sold the toys received over the years, 
// each for P levs and added the amount to the money saved. With the money she wanted to buy a washing machine for X leva. 
// Write a program to calculate how much money she has raised and whether she has enough to buy a washing machine.

package main

import (
	"fmt"
	"math"
)

func main() {
	var age, priceToy int
	fmt.Scanln(&age)
	var priceWashingMachine float64
	fmt.Scanln(&priceWashingMachine)
	fmt.Scanln(&priceToy)

	var savedMoney float64 = 0
	var countJoys int = 0

	for birthday := 1; birthday <= age; birthday++ {
		if birthday % 2 == 0 {
			savedMoney += float64(birthday) * 5
			savedMoney--
		} else {
			countJoys++
		}
	}
	
	savedMoney += (float64(countJoys) * float64(priceToy))
	var diff float64 = math.Abs(savedMoney - priceWashingMachine)
	if savedMoney >= priceWashingMachine {
		fmt.Printf("Yes! %.2f", diff)
	} else {
		fmt.Printf("No! %.2f", diff)
	}
}
