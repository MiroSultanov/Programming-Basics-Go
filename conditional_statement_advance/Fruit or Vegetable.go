// Write a program that reads a product name entered by the user and checks if it is a fruit or vegetable.
// • "Fruits" are banana, apple, kiwi, cherry, lemon and grapes
// • Vegetables are tomato, cucumber, pepper and carrot
// • All others are "unknown"
// Display "fruit", "vegetable" or "unknown" according to the entered product.

package main

import "fmt"

func main(){
	var product string
	fmt.Scanln(&product)

	switch product{
	case "banana", "apple", "kiwi", "cherry", "lemon", "grapes":
		fmt.Print("fruit")
	case "tomato", "cucumber", "pepper", "carrot":
		fmt.Print("vegetable")
	default:
		fmt.Print("unknown")
	}

}
