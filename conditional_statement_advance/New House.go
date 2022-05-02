// Marin and Nelly are buying a house not far from Sofia. Nelly loves flowers so much that she convinces you to write a program that will calculate how much it will cost 
// them, to plant a certain number of flowers and whether the available budget will be enough. Different flowers have different prices.
// flower                                  Rose    Dahlia    Tulip    Narcissus    Gladiolus
// Price per piece in BGN                    5      3.80      2.80      3             2.50
// There are the following discounts:
// • If Nelly buys more than 80 Roses - 10% discount from the final price
// • If Nelly buys more than 90 Dahlias - 15% discount from the final price
// • If Nelly buys more than 80 Tulips - 15% discount from the final price
// • If Nelly buys less than 120 Narcissus - the price increases by 15%
// • If Nelly Buys less than 80 Gladiolus - the price increases by 20%

package main

import "fmt"

func main(){
	var typeFlower string
	fmt.Scanln(&typeFlower)

	var numberOfFlower int 
	fmt.Scanln(&numberOfFlower)

	var budget int 
	fmt.Scanln(&budget)

	var pricePerFlower float32

	if typeFlower == "Roses"{
		if numberOfFlower > 80 {
			pricePerFlower = (5.0 * float32(numberOfFlower)) - (5.0 * float32(numberOfFlower) * 0.1)
		}else{
			pricePerFlower = 5.0 * float32(numberOfFlower)
		}
	}else if typeFlower == "Dahlias"{
		if numberOfFlower > 90 {
			pricePerFlower = (3.80 * float32(numberOfFlower)) - (3.80 * float32(numberOfFlower) * 0.15)
		}else{
			pricePerFlower = 3.80 * float32(numberOfFlower)
		}
	}else if typeFlower == "Tulips" {
		if numberOfFlower > 80 {
			pricePerFlower = (2.80 * float32(numberOfFlower)) - (2.80 * float32(numberOfFlower) * 0.15)
		}else{
			pricePerFlower = 2.80 * float32(numberOfFlower)
		}
	}else if typeFlower == "Narcissus" {
		if numberOfFlower > 120 {
			pricePerFlower = (3.0 * float32(numberOfFlower)) - (3.0 * float32(numberOfFlower) * 0.15)
		}else{
			pricePerFlower = 3.0 * float32(numberOfFlower)
		}
	}else if typeFlower == "Gladiolus" {
		if numberOfFlower > 80{
			pricePerFlower = (2.50 * float32(numberOfFlower)) - (2.50 * float32(numberOfFlower) * 0.20)
		}else{
			pricePerFlower = 2.50 * float32(numberOfFlower)
		}
	}

	if float32(budget) > pricePerFlower {
		var leftMoney = float32(budget) - pricePerFlower
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", numberOfFlower, typeFlower, leftMoney)
	}else{
		var neededMoney = pricePerFlower - float32(budget)
		fmt.Printf("Not enough money, you need %.2f leva more.", neededMoney)
	}
}
