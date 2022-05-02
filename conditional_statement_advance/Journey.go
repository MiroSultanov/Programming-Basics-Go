package main

import "fmt"

func main(){
	var budget float64
	fmt.Scanln(&budget)

	var season string
	fmt.Scanln(&season)

	var destination string = ""
	var place string = ""
	var price float64 = 0

	if budget <= 100 {
		destination = "Bulgaria"
		if season == "summer" {
			place = "Camp"
			price = budget * 0.3
		}else if season == "winter" {
			place = "Hotel"
			price = budget * 0.7
		}
	}else if budget <= 1000 {
		destination = "Balkans"
		if season == "summer" {
			place = "Camp"
			price = budget * 0.4
		}else if season == "winter" {
			place = "Hotel"
			price = budget * 0.8
		}
	}else if budget > 1000 {
		destination = "Europe"
		if season == "summer" {
			place = "Hotel"
			price = budget * 0.9
		}else if season == "winter" {
			place = "Hotel"
			price = budget * 0.9
		}
	}
	fmt.Printf("Somewhere in %s\n", destination)
	fmt.Printf("%s - %.2f", place, price)
}