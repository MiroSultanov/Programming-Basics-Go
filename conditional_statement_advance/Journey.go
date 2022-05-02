// Strange, but most people plan their vacation early. A young programmer has a certain budget and free time in a given season. Write a program that will accept
// the budget and the season at the entrance, and at the exit to earn, where the programmer will rest and how much he will spend.
// The budget determines the destination, and the season determines how much of the budget you will spend. If it is summer he will rest at the campsite and in the 
// winter at a hotel. If he is in Europe, he will stay in a hotel, regardless of the season. Each campsite or hotel, according to the destination, has its own price 
// that corresponds to a certain percentage of the budget:
// • At BGN 100. or less - somewhere in Bulgaria
// o Summer - 30% of the budget
// o Winter - 70% of the budget
// • At BGN 1,000. or a little - somewhere in the Balkans
// o Summer - 40% of the budget
// o Winter - 80% of the budget
// • For more than BGN 1,000. - somewhere in Europe
// o When traveling in Europe, regardless of the season will spend 90% of the budget.

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
