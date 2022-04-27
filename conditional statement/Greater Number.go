package main

import "fmt"

func main(){
	var firstNumber int
	fmt.Scanln(&firstNumber)
	var secondNumber int
	fmt.Scanln(&secondNumber)

	if firstNumber >= secondNumber{
		fmt.Println(firstNumber)
	}else{
		fmt.Println(secondNumber)
	}
}