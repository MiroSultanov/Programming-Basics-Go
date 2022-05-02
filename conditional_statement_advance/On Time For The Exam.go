package main

import "fmt"

func main () {

	var examHour, examMiuntes int
	fmt.Scanln(&examHour)
	fmt.Scanln(&examMiuntes)

	var arriveHour, arriveMinute int
	fmt.Scanln(&arriveHour)
	fmt.Scanln(&arriveMinute)

	var examInMinutes int = examHour * 60 + examMiuntes
	var arriveInMinutes int = arriveHour * 60 + arriveMinute

	if arriveInMinutes > examInMinutes {
		fmt.Println("Late")
		var late int = arriveInMinutes - examInMinutes
		if late < 60 {
			fmt.Printf("%d minutes after the start", late)
		} else {
			var lateHour int = late / 60
			var lateMinutes int = late % 60
			fmt.Printf("%d:%02d hours after the start", lateHour, lateMinutes)
		}
	} else if arriveInMinutes == examInMinutes || examInMinutes - arriveInMinutes <= 30 {
		fmt.Println("On time")
		if examInMinutes - arriveInMinutes <= 30 {
			fmt.Printf("%d minutes before the start", examInMinutes - arriveInMinutes)
		}
	} else if examInMinutes - arriveInMinutes > 30 {
		fmt.Println("Early")
		var early int = examInMinutes - arriveInMinutes
		if early < 60 {
			fmt.Printf("%d minutes before the start", early)
		} else {
			var earlyHour int = early / 60
			var earlyMinutes int = early % 60
			fmt.Printf("%d:%02d hours before the start", earlyHour, earlyMinutes)
		}
	}



}