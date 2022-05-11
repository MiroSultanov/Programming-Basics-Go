package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for number := 1; number < n; number = number * 2 + 1 {
		fmt.Println(number)
	}
}