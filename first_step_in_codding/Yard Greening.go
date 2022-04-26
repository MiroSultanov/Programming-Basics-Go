// Bozhidara has several houses on the Black Sea coast and wants to green the yards of some of them, thus creating a cozy atmosphere and comfort for its guests.
// She has hired a company for this purpose.
// Write a program that calculates the amount needed that Bozhidara will have to pay to the project contractor. The price per square meter is BGN 7.61 with VAT. 
// Because her yard is quite large, the contractor company offers an 18% discount on the final price.

// Input
// Only one line is read from the console:
// 1. Sq. meters to be landscaped - real number in the range [0.00… 10000.00]

// Output
// Two lines are printed on the console:
// • "The final price is: {final price of the service} lv."
// • "The discount is: {discount} lv."

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
