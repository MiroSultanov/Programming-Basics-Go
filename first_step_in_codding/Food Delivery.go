// The restaurant opens its doors and offers several menus at preferential prices:
// • Chicken menu - BGN 10.35.
// • Fish menu - BGN 12.40.
// • Vegetarian menu - BGN 8.15.
// Write a program that calculates how much it will cost a group of people to order food to take home.
// The group will also order a dessert, the price of which is equal to 20% of the total bill (excluding delivery).
// The delivery price is BGN 2.50 and is finally charged.

// Input
// 3 lines are read from the console:
// • Number of chicken menus - integer in the range [0… 99]
// • Number of fish menus - integer in the range [0… 99]
// • Number of vegetarian menus - integer in the range [0… 99]

// Output
// Print one line on the console: "{order price}"

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
