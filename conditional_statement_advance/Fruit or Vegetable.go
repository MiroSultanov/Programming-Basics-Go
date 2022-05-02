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