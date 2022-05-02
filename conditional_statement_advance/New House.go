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