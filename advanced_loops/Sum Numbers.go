// # Write a program that reads an integer from the console and integers on each line until their sum is greater than or equal to the original number. 
// # At the end of the reading, print the sum of the entered numbers.

package main

import "fmt"

func main() {
	var firstNumber int
	fmt.Scanln(&firstNumber)

	var sum int = 0

	for sum < firstNumber {
		var number int
		fmt.Scanln(&number)
		sum += number
	}

	fmt.Print(sum)
}
