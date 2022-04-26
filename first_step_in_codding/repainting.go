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