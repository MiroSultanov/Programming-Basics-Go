package main

import "fmt"

func main(){
	var tripPrice float32
	fmt.Scanln(&tripPrice)
	var puzzelCount int
	fmt.Scanln(&puzzelCount)
	var dollCount int
	fmt.Scanln(&dollCount)
	var bearCount int
	fmt.Scanln(&bearCount)
	var minionsCount int
	fmt.Scanln(&minionsCount)
	var trucksCount int
	fmt.Scanln(&trucksCount)

	var profit float32 = (float32(puzzelCount) * 2.60) + (float32(dollCount) * 3.00) + (float32(bearCount) * 4.10) + (float32(minionsCount) * 8.20) + (float32(trucksCount) * 2.00)
	var sumToys int = puzzelCount + dollCount + bearCount + minionsCount + trucksCount
	
	if sumToys >= 50{
		profit = 0.75 * profit
	}

	var rent float32 = 0.1 * profit
	var totalSum float32 = profit - rent

	if totalSum >= tripPrice{
		var leftMoney float32 = totalSum - tripPrice
		fmt.Printf("Yes! %.2f lv left.", leftMoney)
	}else{
		var neededMoney float32 = tripPrice - totalSum
		fmt.Printf("Not enough money! %.2f lv needed.", neededMoney)
	}
}