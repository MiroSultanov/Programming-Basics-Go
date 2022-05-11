// # Write a program that calculates how much money is in the account after you make a certain number of installments. 
// # On each line you will receive the amount you need to deposit in the account until you receive the command "NoMoreMoney". 
// # For each amount received, the console must display "Increase:" + the amount and add it to the account. If you get a number less than 0, "Invalid operation!" 
// # Must be displayed on the console. and the program ends. 
// # When the program ends, "Total:" must be printed + the total amount in the account formatted to the second decimal place.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	var bankAccount float32 = 0

	for input != "NoMoreMoney" {
		sum, _ := strconv.ParseFloat(input, 32)

		if sum > 0 {
			bankAccount += float32(sum)
			fmt.Printf("Increase: %.2f\n",sum)
		}else{
			fmt.Println("Invalid operation!")
			break
		}

		fmt.Scanln(&input)
	} 

	fmt.Printf("Total: %.2f", bankAccount)
}
