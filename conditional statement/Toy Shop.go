// Petya has a toy store. She receives a large order that she must fulfill. With the money he will earn he wants to go on a trip.
// Toy prices:
// • Puzzle - BGN 2.60.
// • Talking doll - BGN 3
// • Teddy bear - BGN 4.10.
// • Minion - BGN 8.20.
// • Truck - BGN 2
// If the ordered toys are 50 or more, the store makes a 25% discount on the total price. Petya has to give 10% of the earned money for the rent of the store. 
// To calculate whether the money will be enough for her to go on a trip.

// Input
// 6 lines are read from the console:
// 1. Price of the trip - real number in the interval [1.00… 10000.00]
// 2. Number of puzzles - integer in the interval [0… 1000]
// 3. Number of talking dolls - integer in the interval [0… 1000]
// 4. Number of teddy bears - integer in the range [0… 1000]
// 5. Number of minions - integer in the interval [0… 1000]
// 6. Number of trucks - integer in the range [0… 1000]

// Output
// The following is printed on the console:
// • If the money is enough it is printed:
// o "Yes! {remaining money} lv left."
// • If the money is NOT enough it is printed:
// o "Not enough money! lv needed."

// The result must be formatted to the second decimal place.

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
