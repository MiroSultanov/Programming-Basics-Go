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