// To write a program that reads the day of the week (text) - entered by the user and prints on the console the price of a movie ticket according to the day of the week:
// Monday Tuesday Wednesday Thursday Friday  Saturday Sunday
//  12       12       14       14       12      16      16

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
