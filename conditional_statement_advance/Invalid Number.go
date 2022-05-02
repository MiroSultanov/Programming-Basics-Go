// A number is valid if it is in the range [100… 200] or is 0. Write a program that reads an integer entered by the user and print "invalid" 
// if the number entered is not valid.

package main

import "fmt"

func main(){
	var number int
	fmt.Scanln(&number)

	var isValid bool = (number >= 100 && number <= 200) || (number == 0)

	if !isValid{
		fmt.Print("invalid")	
	}
}
