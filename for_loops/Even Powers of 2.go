// Write a program that reads the number n entered by the user and prints the even degrees 2 ≤ 2n: 20, 22, 24, 26,…, 2n.

package main

import "fmt"
import "math"

func main(){
	var n int 
	fmt.Scanln(&n)

	for step := 0; step <= n; step += 2 {
		fmt.Println(math.Pow(2, float64(step)))
	}
}
