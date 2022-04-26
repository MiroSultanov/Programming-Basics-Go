// Write a program that calculates the cost of buying dog and cat food. The food is bought from a pet store, as one package of dog food costs BGN 2.50,
// and a package of cat food costs BGN 4.

// Input
// 2 lines are read from the console:
// 1. Number of packages of dog food - integer in the range [0… 100]
// 2. Number of packages of cat food - integer in the range [0… 100]

// Output
// The following is printed on the console:
// "{final amount} lv."

package main

import "fmt"

func main(){
	var numberDogfood float32
	fmt.Scanln(&numberDogfood)
	var numberCatfood float32
	fmt.Scanln(&numberCatfood)
	var totalDogfood float32 = numberDogfood * 2.50
	var totalCatfood float32 = numberCatfood * 4
	var totalSum float32= totalCatfood + totalDogfood

	fmt.Printf("%.1f lv." ,totalSum)
}
