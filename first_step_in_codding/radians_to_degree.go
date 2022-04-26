package main

import "fmt"
import "math"

func main(){
	var radians float32
	fmt.Scanln(&radians)
	
	var degree float32 = radians * 180 / math.Pi

	fmt.Println(degree)
}