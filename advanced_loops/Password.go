package main

import "fmt"

func main() {
	var username string 
	fmt.Scanln(&username)

	var password string
	fmt.Scanln(&password)

	var enterPass string
	fmt.Scanln(&enterPass)

	for enterPass != password {
		fmt.Scanln(&enterPass)
	}

	fmt.Printf("Welcome %s!", username)

}