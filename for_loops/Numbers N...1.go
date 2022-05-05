// Write a program that reads a positive integer n entered by the user and prints the numbers n to 1 in reverse order. 
// The number n entered will always be greater than 1.

package main

import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)

	for number := n; number >= 1; number -- {
		fmt.Println(number)
	}
}
