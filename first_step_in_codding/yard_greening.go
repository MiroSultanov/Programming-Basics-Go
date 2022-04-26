package main

import "fmt"

func main(){
	var greeningYard float32
	fmt.Scanln(&greeningYard)
	var pricePerWholeYard float32 = greeningYard * 7.61
	var discount float32 = 0.18 * pricePerWholeYard
	var totalPrice float32 = pricePerWholeYard - discount

	fmt.Printf("The final price is: %.2f lv.\n", totalPrice)
	fmt.Printf("The discount is: %.2f lv.", discount)

}