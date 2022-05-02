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