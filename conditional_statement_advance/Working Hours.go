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