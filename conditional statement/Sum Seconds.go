// Three athletes finish in a matter of seconds (between 1 and 50). Write a program that reads the times of the competitors in seconds entered by the user and 
// calculates their total time in the format "minutes: seconds". Display the seconds with leading zero (2 -> "02", 7 -> "07", 35 -> "35").

package main

import "fmt"

func main(){
	var firstTime int
	fmt.Scanln(&firstTime)
	var seconTime int
	fmt.Scanln(&seconTime)
	var thirdTime int
	fmt.Scanln(&thirdTime)

	var totalTime int = firstTime + seconTime + thirdTime
	var minutes int = totalTime / 60
	var seconds int = totalTime % 60

	if seconds < 10{
		fmt.Printf("%d:0%d", minutes , seconds)
	}else{
		fmt.Printf("%d:%d", minutes, seconds)
	}
}
