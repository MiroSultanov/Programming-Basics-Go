// Write a program that calculates how much you will receive at the end of the deposit period at a certain interest rate. Use the following formula:
// amount = deposited amount + term of the deposit * (deposited amount * annual interest rate) / 12)

// Input
// 3 lines are read from the console:
// 1. Deposit amount - real number in the interval [100.00… 10000.00]
// 2. Term of the deposit (in months) - integer in the interval [1… 12]
// 3. Annual interest rate - real number in the interval [0.00… 100.00]

// Output
// Print the amount on the console at the end of the term.

package main

import "fmt"

func main(){
	var depositSum float32
	fmt.Scanln(&depositSum)

	var month int 
	fmt.Scanln(&month)

	var percent float32
	fmt.Scanln(&percent)

	var interest float32 = depositSum * (percent / 100.0)
	var interetPerMonth float32 = interest / 12.0
	var totalSum float32 = depositSum + float32(month) * interetPerMonth

	fmt.Println(totalSum)
}
