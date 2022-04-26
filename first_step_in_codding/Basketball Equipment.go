// Jesse decides he wants to play basketball, but he needs equipment to train. Write a program that calculates what costs Jesse will have if he starts training,
// knowing how much the basketball training fee is for a period of 1 year. Required equipment:
// • Basketball sneakers - their price is 40% less than the fee for one year
// • Basketball team - its price is 20% cheaper than sneakers
// • Basketball - its price is 1/4 of the price of the basketball team
// • Basketball accessories - their price is 1/5 of the price of a basketball

// Input
// 1 line is read from the console:
// • The annual fee for basketball training - an integer in the interval [0… 9999]

// Output
// To print on the console how much Jesse will spend if he starts playing basketball.

package main

import "fmt"

func main() {
	var tax int
	fmt.Scanln(&tax)

	var sneakers float32 = float32(tax) - (float32(tax) * 0.4)
	var equipment float32 = float32(sneakers) - (float32(sneakers) * 0.2)
	var ball float32 = float32(equipment) * 0.25
	var accessories float32 = float32(ball) / 5

	var totalsum float32 = float32(tax) + float32(sneakers) + float32(equipment) + float32(ball) + float32(accessories)

	fmt.Println(totalsum)
}
