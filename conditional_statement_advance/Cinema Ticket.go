package main

import "fmt"

func main(){
	var day string 
	fmt.Scanln(&day)

	switch day{
	case "Monday", "Tuesday", "Friday":
		fmt.Print(12)
	case "Wednesday", "Thursday":
		fmt.Print(14)
	case "Saturday", "Sunday":
		fmt.Print(16)
	default:
		fmt.Print("Error")
	}
}