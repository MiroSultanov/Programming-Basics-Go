// Write a program that reads an integer entered by the user and prints whether it is even or odd.
// If even, print "even", if odd, "odd".

package main

import "fmt"

func main(){
	var number int
	fmt.Scanln(&number)

	if number % 2 == 0{
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}
}
