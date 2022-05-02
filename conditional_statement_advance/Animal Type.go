// Write a program that prints the class of the animal according to its name entered by the user.
// • dog -> mammal
// • crocodile, tortoise, snake -> reptile
// • others -> unknown

package main

import "fmt"

func main(){
	var animal string
	fmt.Scanln(&animal)

	switch animal{
	case "dog":
		fmt.Print("mammal")
	case "crocodile", "tortoise", "snake":
		fmt.Print("reptile")
	default:
		fmt.Print("unknown")
	}
}
