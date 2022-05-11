// # Write a program that reads the number n entered by the user and prints all the numbers ≤ n from the sequence: 1, 3, 7, 15, 31,…. 
// # Each subsequent number is calculated by multiplying the previous one by 2 and adding 1.

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for number := 1; number < n; number = number * 2 + 1 {
		fmt.Println(number)
	}
}
