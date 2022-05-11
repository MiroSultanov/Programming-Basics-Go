// # Jesse has decided to raise money for a trip and asks you to help her find out if she will be able to raise the necessary amount. 
// # She saves or spends some of her money every day. If she wants to spend more than her cash, she will spend as much as she has and she will have 0 leva left.

// # Input
// # From the console read:
// # ⦁ Money needed for the trip - real number;
// # ⦁ Cash - real number.
// # Then two lines are read repeatedly:
// # ⦁ Type of action - text with two options: "spend" or "save";
// # ⦁ The amount you will save / spend - a real number.

// # Output
// # The program must be completed in the following cases:
// # ⦁ If Jesse only spends 5 consecutive days, the console should show:
// # ⦁ "You can't save the money."
// # ⦁ "{Total days passed}"
// # ⦁ If Jesse collects the money for the holiday, the console says:
// # ⦁ "You saved the money for {total days passed} days."


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
