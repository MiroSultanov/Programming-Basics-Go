// Summer is very changeable and Victor needs your help. Write a program that, depending on the time of day and the degrees, can recommend Victor what clothes to wear.
// Your friend has different plans for each stage of the day, which require a different look, you can see them in the table.
// Exactly two lines are read from the console:
// • Degrees - an integer in the range [10… 42]
// • Text, time of day - with options - "Morning", "Afternoon", "Evening"
// Time of day / degrees
// 10 <= degrees <= 18           Morning                        Afternoon                       Evening
//                           Outfit = Sweatshirt                Outfit = Shirt                  Outfit = Shirt
//                           Shoes = Sneakers                   Shoes = Moccasins               Shoes = Moccasins

// 18 <degrees <= 24         Outfit¬ = Shirt                    Outfit = T-Shirt                Outfit = Shirt
//                           Shoes = Moccasins                  Shoes = Sandals                 Shoes = Moccasins
 
 
// degrees> = 25            Outfit = T-Shirt                    Outfit = Swim Suit              Outfit = Shirt
//                          Shoes = Sandals                     Shoes = Barefoot                Shoes = Moccasins
 
// To be printed on the console in one line: "It's {degrees} degrees, get your {clothing} and {shoes}."

package main

import "fmt"

func main(){
	var degrees int 
	fmt.Scanln(&degrees)

	var partOfDay string
	fmt.Scanln(&partOfDay)

	var outfit string = ""
	var shoes string = ""

	if degrees >= 10 && degrees <=18 {
		switch partOfDay{
		case "Morning":
			outfit = "Sweatshirt"
			shoes = "Sneakers"
		case "Afternoon","Evening":
			outfit = "Shirt"
			shoes = "Moccasin" 
		}
	}else if degrees > 18 && degrees <= 24{
		switch partOfDay{
		case "Morning":
			outfit = "Shirt"
			shoes = "Moccasins"
		case "Afternoon":
			outfit = "T-Shirt"
			shoes = "Sandals"
		case "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
		}
	}else if degrees >= 25{
		switch partOfDay{
		case "Morning":
			outfit = "T-Shirt"
			shoes = "Sandals"
		case "Afternoon":
			outfit = "Swim Suit"
			shoes = "Barefoot"
		case "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
	}
}

	fmt.Printf("It's %d degrees, get your %s and %s.", degrees ,outfit, shoes)
}
