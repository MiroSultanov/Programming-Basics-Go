package main

import "fmt"

func main(){
	var numberDogfood float32
	fmt.Scanln(&numberDogfood)
	var numberCatfood float32
	fmt.Scanln(&numberCatfood)
	var totalDogfood float32 = numberDogfood * 2.50
	var totalCatfood float32 = numberCatfood * 4
	var totalSum float32= totalCatfood + totalDogfood

	fmt.Printf("%.1f lv." ,totalSum)
}