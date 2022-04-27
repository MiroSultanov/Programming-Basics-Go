// Filming for the long-awaited film "Godzilla vs. Kong" begins. Screenwriter Adam Wingard asks you to write a program to calculate whether the funds provided are
// enough to shoot the film. The photos will require a certain number of extras, clothing for each extra and decor.
// It is known that:
// • The set for the film is worth 10% of the budget.
// • For more than 150 extras, there is a 10% discount on clothing.

// Input
// 3 lines are read from the console:
// Row 1. Movie budget - real number in the interval [1.00… 1000000.00]
// Row 2. Number of extras - integer in the range [1… 500]
// Row 3. Price for clothing of one extra - real number in the interval [1.00… 1000.00]

// Output
// Two lines must be printed on the console:
// • If the money for decor and clothes is more than the budget:
// o "Not enough money!"
// o "Wingard needs {money not enough for the movie} leva more."
// • If the money for decor and clothes is less than or equal to the budget:
// about "Action!"
// o "Wingard starts filming with {remaining money} left left."
// The result must be formatted to the second decimal place.

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
