// # Write a program that reads integers from the console until a "stop" command is received. Find the sum of all entered primes and the sum of all
// # entered primes. Since, by definition in mathematics, negative numbers cannot be prime, if a negative number is input, the following message
// # "Number is negative." In this case, the entered number is ignored and is not added to either of the two amounts, and the program continues its execution,
// # waiting for the next number to be entered.
// # At the output to print in two lines the two amounts found in the following format:
// #  "Sum of all prime numbers is: {prime numbers sum}"
// # ⦁ "Sum of all non prime numbers is: {nonprime numbers sum}"

package main

import (
	"fmt"
	"strconv"
)

func main4() {

	var input string
	fmt.Scanln(&input)

	var sumPrime, sumNonPrime int = 0, 0
	
	
	for input != "stop" {
		
		number, _ := strconv.Atoi(input) 

		if number < 0 {
			fmt.Println("Number is negative.")
		} else {
			
			var count int = 0 
			
			for i := 1; i <= number; i++ {
				if number % i == 0 {
					count++
				}	
			}
			
			if count == 2 {
				sumPrime += number
			} else {
				
				sumNonPrime += number
			}
		}
		fmt.Scanln(&input) 
	}

	fmt.Printf("Sum of all prime numbers is: %d\n", sumPrime)
	fmt.Printf("Sum of all non prime numbers is: %d", sumNonPrime)


}
