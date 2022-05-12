// # Write a program that reads two six-digit integers from the console. The first number entered will always be less than the second. 
// # On the console to be printed on 1, line separated by a space, all numbers that are between the two numbers read by the console and meet the condition
// # that the sum of the digits of even and odd positions are equal. 
// # If there are no numbers that meet the console condition, no result is displayed.

package main


import "fmt"

func main1() {
	var firstNumber, secondNumber int
	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)

	for number := firstNumber; number <= secondNumber; number++ {
		var units int = number % 10

		var tens int = number / 10 % 10
	
		var hundreds int = number / 100 % 10
		
		var thousands int = number / 1000 % 10
		
		var tenThoudsands int = number / 10000 % 10
		
		var hundredThousands int = number / 100000

		
		if units + hundreds + tenThoudsands == tens + thousands + hundredThousands {
			fmt.Printf("%d ", number)
		}
	}
}
