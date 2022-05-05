package main

import "fmt"
import "math"

func main(){
	var n int 
	fmt.Scanln(&n)

	var max int = math.MinInt16
	var sum int = 0

	for number := 1; number <= n; number ++ {
		var value int 
		fmt.Scanln(&value)

		if value > max{
			max = value
		}

		sum = sum + value
	}

	var sumOthers int = sum - max

	if max == sumOthers {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", max)
	}else {
		fmt.Println("No")
		var diff float64 = math.Abs(float64(max - sumOthers))
		fmt.Printf("Diff = %.0f", diff) 
	}
}