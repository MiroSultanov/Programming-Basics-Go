// Tony and his friends loved to go fishing, they were so passionate about fishing that they decided to go fishing by boat. 
// The price of renting a boat depends on the season and the number of fishermen.
// The price depends on the season:
// • The price for renting the ship in the spring is 3000 BGN.
// • The price for renting the ship in summer and autumn is 4200 BGN.
// • The price for renting the ship in winter is 2600 BGN.
// Depending on its number, the group enjoys a discount:
// • If the group is up to 6 people inclusive - 10% discount.
// • If the group is from 7 to 11 people inclusive - 15% discount.
// • If the group is from 12 upwards - 25% discount.
// Fishermen enjoy an additional 5% discount if they are an even number unless it is autumn - then they do not have an additional discount.
// Write a program to calculate whether fishermen will raise enough money.

package main

import "fmt"

func main (){
	var budget int 
	fmt.Scanln(&budget)

	var season string
	fmt.Scanln(&season)

	var numberFisherman int 
	fmt.Scanln(&numberFisherman)

	var boatRent float32 = 0
	
	if season == "Spring"{
		if numberFisherman <= 6 {
			boatRent = 3000 - (3000 * 0.10)
		}else if numberFisherman >= 7 || numberFisherman <= 11 {
			boatRent = 3000 - (3000 * 0.15)
		}else if numberFisherman > 12 {
			boatRent = 3000 - (3000 * 0.25)
		}

	}else if season == "Summer" || season == "Autumn"{
		if numberFisherman <= 6 {
			boatRent = 4200 - (4200 * 0.10)
		}else if numberFisherman >= 7 || numberFisherman <= 11 {
			boatRent = 4200 - (4200 * 0.15)
		}else if numberFisherman > 12 {
			boatRent = 4200 - (4200 * 0.25)
		}
		
	}else if season == "Winter"{
		if numberFisherman <= 6 {
			boatRent = 2600 - (2600 * 0.10)
		}else if numberFisherman >= 7 || numberFisherman <= 11 {
			boatRent = 2600 - (2600 * 0.15)
		}else if numberFisherman > 12 {
			boatRent = 2600 - (2600 * 0.25)
		}
	}	

	if numberFisherman % 2 == 0 && season != "Autumn"{
		boatRent = boatRent - (boatRent * 0.05)
	}

	if float32(budget) >= boatRent {
		var leftMoney float32 = float32(budget) - boatRent
		fmt.Printf("Yes! You have %.2f leva left.", leftMoney)
	}else{
		var neededMoney float32 = boatRent - float32(budget)
		fmt.Printf("Not enough money! You need %.2f leva.", neededMoney)
	}			
}
