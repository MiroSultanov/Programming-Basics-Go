// Write a console program that enters two integers (the sides of rectangles a and b) and calculates the face of a rectangle with these sides.

package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	var area int = a * b
	fmt.Println(area)
}
