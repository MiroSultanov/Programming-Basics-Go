// Write a program that reads an hour of the day (integer) and a day of the week (text) - entered by the user and checks whether the company's office is open, 
// the office hours are from 10 am to 6 pm, Monday to Saturday including

package main

import "fmt"

func main(){
	var hour int
	fmt.Scanln(&hour)

	var dayOfWeek string
	fmt.Scanln(&dayOfWeek)

	if hour < 10 || hour > 18 || dayOfWeek == "Sunday"{
		fmt.Print("closed")
	}else{
		fmt.Print("open")
	}
}
