package main

import "fmt"

func main(){
	var chickenMenu int
	fmt.Scanln(&chickenMenu)
	var fishMenu int
	fmt.Scanln(&fishMenu)
	var veganMenu int
	fmt.Scanln(&veganMenu)

	var priceChickenMenu float32 = float32(chickenMenu) * 10.35
	var priceFishMenu float32 = float32(fishMenu) * 12.40
	var priceVegenMenu float32 = float32(veganMenu) * 8.15

	var totalPrice float32 = priceChickenMenu + priceFishMenu + priceVegenMenu
	var priceDesert float32 = float32(totalPrice) * 0.2
	var delivery float32 = 2.50

	var finalPrice float32 = float32(totalPrice) + float32(priceDesert) + float32(delivery)

	fmt.Printf("%.2f", finalPrice)
}