package main

import "fmt"

func main(){
	var lenght int
	fmt.Scanln(&lenght)
	var width int
	fmt.Scanln(&width)
	var height int
	fmt.Scanln(&height)
	var percent float32
	fmt.Scanln(&percent)

	var volumeTank int = lenght * width * height
	var volumeLitter float32 = float32(volumeTank) * 0.001 
	var busy float32 = percent / 100
	var neededLitters float32 = float32(volumeLitter) *(1.0 - busy)
	fmt.Println(neededLitters)
}