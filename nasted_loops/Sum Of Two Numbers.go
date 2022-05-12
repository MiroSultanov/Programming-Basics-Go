package main

import (
	"fmt"
)

func main() {
	var startNumber int 
	fmt.Scanln(&startNumber)

	var endNumber int
	fmt.Scanln(&endNumber)

	var magicNumber int 
	fmt.Scanln(&magicNumber)

	var count int = 0
	var isFound bool = false

	for first := startNumber; first <= endNumber; first++ {
		for second := startNumber; second <= endNumber; second++ {
			count++

			if first + second == magicNumber{
				fmt.Printf("Combination N:%d (%d + %d = %d)", count, first, second, magicNumber)
				isFound =true
				break
			}
		}
		if isFound {
			break
		}
	}
	if !isFound {
		fmt.Printf("%d combinations - neither equals %d", count, magicNumber)
	}
}