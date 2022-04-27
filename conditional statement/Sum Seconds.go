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