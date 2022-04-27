package main

import "fmt"

func main(){
	var budget float32
	fmt.Scanln(&budget)

	var numberStatist int
	fmt.Scanln(&numberStatist)

	var pricePerDress float32
	fmt.Scanln(&pricePerDress)

	var sumDecor = 0.1 * budget
	var priceClothes float32 = float32(numberStatist) * pricePerDress

	if numberStatist > 150{
		priceClothes = priceClothes - priceClothes * 0.1
	}

	var totaPrice float32 = sumDecor + priceClothes

	if totaPrice <= budget{
		var leftMoney = budget - totaPrice
		fmt.Println("Action!")
		fmt.Printf("Wingard starts filming with %.2f leva left.", leftMoney)
	}else {
		var neededMoney float32 = totaPrice - budget
		fmt.Println("Not enough money!")
		fmt.Printf("Wingard needs %.2f leva more.", neededMoney)
	}
}