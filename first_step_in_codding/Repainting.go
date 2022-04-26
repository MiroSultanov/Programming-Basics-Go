// Rumen wants to repaint the living room and has hired craftsmen for this purpose. Write a program that calculates the cost of repairs, given the following prices:
// • Protective nylon - BGN 1.50 per square meter
// • Paint - BGN 14.50 per liter
// • Thinner for paint - BGN 5.00 per liter
// Just in case, to the necessary materials, Rumen wants to add another 10% of the amount of paint and 2 sq.m. nylon, of course BGN 0.40 for bags. 
// The amount paid to the masters for 1 hour of work is equal to 30% of the sum of all costs for materials.

// Input
// The input is readable from the console and contains exactly 4 lines:
// 1. Required amount of nylon (in sq.m.) - an integer in the range [1 ... 100]
// 2. Required amount of paint (in liters) - an integer in the range [1… 100]
// 3. Quantity of diluent (in liters) - integer in the interval [1… 30]
// 4. The hours for which the masters will do the work - an integer in the interval [1… 9]

// Output
// Print one line on the console:
// • "{sum of all costs}"

package main

import "fmt"

func main(){
	var neededNylon int
	fmt.Scanln(&neededNylon)
	var neededPaint int
	fmt.Scanln(&neededPaint)
	var neededThinner int
	fmt.Scanln(&neededThinner)
	var neededHours int
	fmt.Scanln(&neededHours)

	var sumNylon float32 = (float32(neededNylon) + float32(2)) * 1.50
	var sumPaint float32 = (float32(neededPaint) * 14.50) * 0.1 + (float32(neededPaint) * 14.50)
	var sumThinner float32 = float32(neededThinner) * 5.00
	var sumBags float32 = 0.40

	var totalSum float32 = sumNylon + sumPaint + sumThinner + sumBags
	var sumBuilders float32 = (float32(totalSum) * 0.30) * float32(neededHours)
	var finalSum float32 = float32(totalSum) + float32(sumBuilders)

	fmt.Printf("%.2f", finalSum)
}
