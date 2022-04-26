// For his birthday, Lubomir received an aquarium in the shape of a parallelepiped. Initially, we read from the console in separate rows its 
// dimensions - length, width and height in centimeters. It is necessary to calculate how many liters of water the aquarium will collect if it is known that
// a certain percentage of its capacity is occupied by sand, plants, heater and pump.
// One liter of water is equal to one cubic decimeter (1 l = 1 dm3).
// Write a program that calculates the liters of water needed to fill the aquarium.

// Input
// 4 lines are read from the console:
// 1. Length in cm - integer in the interval [10… 500]
// 2. Width in cm - an integer in the interval [10… 300]
// 3. Height in cm - an integer in the range [10… 200]
// 4. Percentage - real number in the interval [0.000… 100.000]

// Output
// Print a number on the console:
// • liters of water that will collect the aquarium.

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
