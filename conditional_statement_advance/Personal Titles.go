// Write a console program that reads the age (decimal number) and gender ("m" or "f") entered by the user and prints an address from the following:
// • "Mr." - male (sex "m") aged 16 or over
// • "Master" - a boy (gender "m") under 16 years
// • "Ms." - a woman (sex "f") aged 16 or over
// • "Miss" - a girl (gender "f") under 16 years

package main

import "fmt"

func main(){
	var age float32
	fmt.Scanln(&age)

	var gender string
	fmt.Scanln(&gender)

	if gender == "m"{
		if age < 16{
			fmt.Print("Master")
		}else{
			fmt.Print("Mr.")
		}
	}else if gender == "f"{
		if age < 16{
			fmt.Print("Miss")
		}else{
			fmt.Print("Ms.")
		}
	}
}
