// Write a program that reads the number n entered by the user and prints the numbers from 1 to n through 3.

package main

import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)

	for number := 1; number <= n; number += 3 {
		fmt.Println(number)
	}
}
