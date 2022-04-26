package main

import "fmt"

func main(){
	var numberPens int
	fmt.Scanln(&numberPens)
	var numberMarkers int
	fmt.Scanln(&numberMarkers)
	var littersCleaner int
	fmt.Scanln(&littersCleaner)
	var percentDiscount int
	fmt.Scanln(&percentDiscount)

	var pricePerPens float32 = float32(numberPens) * 5.80
	var pricePerMarkers float32 = float32(numberMarkers) * 7.20
	var pricePerCleaner float32 = float32(littersCleaner) * 1.20
	var sum float32 = pricePerPens + pricePerMarkers + pricePerCleaner
	var finalSum float32 = sum - (sum * float32(percentDiscount)/100.0)

	fmt.Println(finalSum)	
}