// Write a program that reads n-number integers submitted by the user and checks that the sum of the numbers of even positions is equal to the
// sum of the numbers of odd positions.
// • If the amounts are equal, print two lines: "Yes" and a new line "Sum =" + the amount;
// • If the amounts are not equal, print two lines: "No" and a new line "Diff =" + the difference.
// The difference is calculated in absolute value.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var sumEven int = 0
	var sumOdd int = 0

	for number := 1; number <= n; number++ {
		var value int
		fmt.Scanln(&value)
		if number % 2 == 0 {
			sumEven += value
		} else {
			sumOdd += value
		}
	}
	var diff int = int(math.Abs(float64(sumEven) - float64(sumOdd)))
	if diff == 0 {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", sumEven)
	} else {
		fmt.Println("No")
		fmt.Printf("Diff = %d", diff)
	}
}
