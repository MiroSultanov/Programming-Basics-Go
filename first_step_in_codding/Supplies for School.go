// The school year has already started and the head of 10B class - Annie has to buy a certain number of packets of chemicals, packets of markers, and a detergent
// for cleaning the board. She is a regular customer of a bookstore, so there is a discount for her, which is a percentage of the total. Write a program that calculates
// how much money Annie will have to raise to pay the bill, keeping in mind the following price list:
// • Package of chemicals - BGN 5.80.
// • Marker package - BGN 7.20.
// • Detergent - BGN 1.20 (per liter)

// Input
// 4 numbers are read from the console:
// • Number of packages of chemicals - an integer in the range [0 ... 100]
// • Number of marker packets - integer in the range [0 ... 100]
// • Liters of board cleaner - integer in the range [0… 50]
// • Percentage reduction - an integer in the range [0 ... 100]

// Output
// Print on the console how much money Annie will need to pay her bill.

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
