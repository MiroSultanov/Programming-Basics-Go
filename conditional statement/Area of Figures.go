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