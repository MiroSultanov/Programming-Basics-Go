// Write a program that reads an integer entered by the user and checks if it is below 100, between 100 and 200 or above 200. If the number is:
// • Under 100 print: "Less than 100"
// • between 100 and 200 print: "Between 100 and 200"
// • over 200 print: "Greater than 200"

package main

import "fmt"

func main(){
	var number int
	fmt.Scanln(&number)

	if number < 100{
		fmt.Println("Less than 100")
	}else if number <= 200{
		fmt.Println("Between 100 and 200")
	}else if number > 200{
		fmt.Println("Greater than 200")
	}
}
