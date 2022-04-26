package main

import "fmt"

func main(){
	var depositSum float32
	fmt.Scanln(&depositSum)

	var month int 
	fmt.Scanln(&month)

	var percent float32
	fmt.Scanln(&percent)

	var interest float32 = depositSum * (percent / 100.0)
	var interetPerMonth float32 = interest / 12.0
	var totalSum float32 = depositSum + float32(month) * interetPerMonth

	fmt.Println(totalSum)
}