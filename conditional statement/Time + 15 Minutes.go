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