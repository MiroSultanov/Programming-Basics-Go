package main

import "fmt"

func main() {
	for number := 7; number <= 1000; number ++ {
		if number % 10 == 7 {
			fmt.Println(number)
		}
	
	}
}