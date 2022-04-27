// Write a console program that reads a rating (integer) entered by the user and prints "Excellent!" if the score is 5 or higher.

package main

import "fmt"

func main(){
	var grade int
	fmt.Scanln(&grade)
	if grade >= 5 {
		fmt.Println("Excellent!")

	}
}
