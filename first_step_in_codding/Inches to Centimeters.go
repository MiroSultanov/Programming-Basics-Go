// Write a program that reads a real number from the console that shows how many inches it is. Convert inches to centimeters 
// by multiplying inches by 2.54 (1 inch = 2.54 centimeters).

package main

import "fmt"

func main() {
	var inches float32
	fmt.Scanln(&inches)
	var cm float32 = inches * 2.54
	fmt.Println(cm)
}
