// Peter wants to buy N video cards, M CPU and P number of RAM. If the number of video cards is greater than that of the processors, it receives a 15% discount on the
// final bill. The following prices apply:
// • Video card - BGN 250 / pc.
// • Processor - 35% of the price of the purchased video cards / pc.
// • Memory frame - 10% of the price of the purchased video cards / pc.
// Calculate the amount needed to purchase the materials and calculate whether the budget will be enough.

// Input
// The entrance consists of four rows:
// 1. Peter's budget - a real number in the interval [0.00… 100000.00]
// 2. The number of video cards - an integer in the range [0… 100]
// 3. The number of processors - an integer in the range [0… 100]
// 4. The number of memory frames - integer in the interval [0… 100]

// Output
// 1 line is printed on the console, which should look like this:
// • If the budget is sufficient:
// "You have {remaining budget} leva left!"
// • If the amount exceeds the budget:
// "Not enough money! You need {necessary amount} left more!"
// Format the result to the second decimal place.

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
