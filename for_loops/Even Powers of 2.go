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