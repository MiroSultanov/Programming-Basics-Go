// In a movie theater, the chairs are arranged in a rectangular shape in r rows and c columns. There are three types of screenings with tickets at different prices:
// • Premiere - premiere screening, at a price of BGN 12.00.
// • Normal - standard screening, at a price of BGN 7.50.
// • Discount - screening for children, pupils and students at a reduced price of BGN 5.00.
// Write a program that reads the type of projection (string), number of rows and number of columns in the hall (integers) entered by the user, and calculates the
// total ticket revenue for a full hall. Print the result in the format as in the examples below, 2 characters after the decimal point.

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
