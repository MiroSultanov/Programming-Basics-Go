// Write a program that reads n-number integers entered by the user and sums them.
// • The number of numbers n is entered from the first line of the input.
// • One integer is entered from the next n lines.
// The program must read the numbers, sum them and print their sum.

package main

import "fmt"

func main(){
	var n int 
	fmt.Scanln(&n)

	var sum int = 0

	for number := 1; number <= n; number ++ {
		var value int 
		fmt. Scanln(&value)

		sum += value
	}
	fmt.Println(sum)
}
