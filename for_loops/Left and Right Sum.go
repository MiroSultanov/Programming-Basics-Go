// Write a program that reads 2 * n numbers of integers submitted by the user and checks if the sum of the first n numbers (left sum) is equal to the sum of the second
// n numbers (right sum). In case of equality, print "Yes, sum =" + the amount; otherwise print "No, diff =" + the difference. The difference is calculated as a 
// positive number (in absolute value).

package main

import "fmt"
import "math"

func main() {
	var n int 
	fmt.Scanln(&n)

	var sumLeft int = 0
	var sumRight int = 0

	for number := 1; number <= n; number++ {
		var value int 
		fmt.Scanln(&value)

		sumLeft += value
	}

	for number := 1; number <= n; number++ {
		var value int
		fmt.Scanln(&value)

		sumRight += value
	}

	if sumLeft == sumRight {
		fmt.Printf("Yes, sum = %d",sumLeft)
	}else{
		var diff int = sumLeft - sumRight
		fmt.Printf("No, diff = %.0f",math.Abs(float64(diff)))
	}
}
