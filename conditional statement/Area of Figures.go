// Write a program in which the user enters the type and dimensions of a geometric figure and calculates its face. The figures are of four types: square, square, 
// rectangle, circle and triangle. The first line of the input reads the type of figure (text with the following options: square, rectangle, circle or triangle).
// • If the figure is a square: the next line reads a fractional number - the length of its side
// • If the figure is a rectangle: on the next two lines read two fractional numbers - the lengths of its sides
// • If the figure is a circle: on the next line reads a fractional number - the radius of the circle
// • If the figure is a triangle: the next two lines read two fractional numbers - the length of its side and the length of the height to it
// Round the result to 3 digits after the decimal point.

package main

import "fmt"
import "math"

func main(){
	var typeFigure string
	fmt.Scanln(&typeFigure)

	if typeFigure == "square"{
		var a float32
		fmt.Scanln(&a)
		var area float32 = a * a
		fmt.Printf("%.3f", area)
	}else if typeFigure == "rectangle"{
		var a float32
		var b float32
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		var area float32 = a * b
		fmt.Printf("%.3f", area)
	}else if typeFigure == "circle"{
		var raduis float32
		fmt.Scanln(&raduis)
		var area float32 = math.Pi * raduis * raduis
		fmt.Printf("%.3f", area)
	}else if typeFigure == "triangle"{
		var a float32
		var h float32
		fmt.Scanln(&a)
		fmt.Scanln(&h)
		var area float32 = (a * h) / 2.0
		fmt.Printf("%.3f", area)
	}
}
