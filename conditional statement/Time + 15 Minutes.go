// Write a program that reads the hour and minutes of the 24-hour day entered by the user and calculates what time it will be in 15 minutes. 
// Print the result in hours: minutes format. The hours are always between 0 and 23, and the minutes are always between 0 and 59. The hours are written in one or two digits.
// Minutes are always displayed in two digits, with a leading zero when necessary.

package main

import "fmt"

func main(){
	var hour int
	fmt.Scanln(&hour)
	var minutes int
	fmt.Scanln(&minutes)

	var currentTimeinMinutes int = hour * 60 + minutes
	var additionalTime int = currentTimeinMinutes + 15

	var finalHour int = additionalTime / 60
	var finalMinutes int = additionalTime % 60

	if finalHour == 24{
		finalHour = 0
	}

	fmt.Printf("%d:%02d", finalHour, finalMinutes)
}
