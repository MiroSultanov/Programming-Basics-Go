// Write a program that reads the day of the week (text), in English - entered by the user. If the day is a working day, it prints on the console - "Working day", 
// if it is a day off - "Weekend". If text other than day of the week is entered, print "Error".

package main

import "fmt"

func main(){
	var dayOfWeek string
	fmt.Scanln(&dayOfWeek)

	switch dayOfWeek{
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Print("Working day")
	case "Saturday", "Sunday":
		fmt.Print("Weekend")
	default:
		fmt.Print("Error")
	}
}
