// Write a program that prints numbers in the range 1 to 1000 that end in 7.

package main

import "fmt"

func main() {
	for number := 7; number <= 1000; number ++ {
		if number % 10 == 7 {
			fmt.Println(number)
		}
	
	}
}
