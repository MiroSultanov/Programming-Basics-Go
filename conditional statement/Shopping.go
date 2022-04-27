package main

import "fmt"

func main(){
	var budget float32
	fmt.Scanln(&budget)

	var countVideoCard int
	fmt.Scanln(&countVideoCard)

	var countCPUs int
	fmt.Scanln(&countCPUs)

	var countRam int
	fmt.Scanln(&countRam)

	var sumVideoCards float32 = float32(countVideoCard) * 250.0
	var sumCPUs float32 = float32(countCPUs) * (0.35 * sumVideoCards)
	var sumRam float32 = float32(countRam) * (0.10 * sumVideoCards)
	var totalSum float32 = sumVideoCards + sumCPUs + sumRam

	if countVideoCard > countCPUs{
		totalSum = totalSum - (0.15 * totalSum)
	}

	if budget >= totalSum{
		var leftMoney float32 = budget - totalSum
		fmt.Printf("You have %.2f leva left!", leftMoney)
	}else{
		var neededMoney = totalSum - budget
		fmt.Printf("Not enough money! You need %.2f leva more!", neededMoney)
	}
}