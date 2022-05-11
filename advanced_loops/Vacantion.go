package main

import "fmt"

func main() {
	var neededMoney float32
	fmt.Scanln(&neededMoney)

	var currentMoney float32
	fmt.Scanln(&currentMoney)

	var countAllDays int
	var countSpendDays int 

	for currentMoney < neededMoney {
		var action string
		fmt.Scanln(&action)

		var sum float32
		fmt.Scanln(&sum)

		countAllDays ++

		if action == "spend" {
			countSpendDays ++ 
			currentMoney -= sum
		}else if action == "save" {
			currentMoney += sum
			countSpendDays = 0
		}

		if currentMoney < 0 {
			currentMoney = 0
		}

		if countSpendDays == 5 {
			fmt.Println("You can't save the money.")
			fmt.Println(countAllDays)
			break
		}
	}

	if currentMoney >= neededMoney {
		fmt.Printf("You saved the money for %d days.", countAllDays)
	} 
}