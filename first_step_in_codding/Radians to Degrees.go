// Write a program that reads an angle in radians (real number) and converts it to degrees. Use the formula: 1 degree = 1 radian * 180 / π. 
// The number π in Go programs is available through math.Pi.

package main

import "fmt"
import "math"

func main(){
	var radians float32
	fmt.Scanln(&radians)
	
	var degree float32 = radians * 180 / math.Pi

	fmt.Println(degree)
}
