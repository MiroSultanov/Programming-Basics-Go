// A student must go to the exam at a certain time. He arrives at the examination room at a certain time of arrival. 
// It is considered that the student arrived on time if he arrived at the time of the exam or up to half an hour before. 
// If he arrived more than 30 minutes earlier, he was early. If he came after the exam time, he was late. Write a program that reads exam time and arrival time 
// and records whether the student arrived on time, was early or late, and by how many hours or minutes was early or late.

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
