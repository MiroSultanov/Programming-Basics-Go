// # Vending machine manufacturers wanted to make their machines return as few change coins as possible. 
// # Write a program that accepts an amount - the rest that needs to be returned and calculates with as few coins as possible.

package main

import (
	"fmt"
)

func main() {

	var change float32 = 0.59
	

	var changeInSt int = int(change * 100) 

	
	var count int = 0 

	for changeInSt != 0 {
		
		if changeInSt >= 200 {
			
			count++
			changeInSt -= 200
		} else if changeInSt >= 100 {
			
			count++
			changeInSt -= 100
		} else if changeInSt >= 50 {
			
			count++
			changeInSt -= 50
		} else if changeInSt >= 20 {
			
			count++
			changeInSt -= 20
		} else if changeInSt >= 10 {
			
			count++
			changeInSt -= 10
		} else if changeInSt >= 5 {
			
			count++
			changeInSt -= 5
		} else if changeInSt >= 2 {
			
			count++
			changeInSt -= 2
		} else if changeInSt >= 1 {
			
			count++
			changeInSt -= 1
		}
	}

	fmt.Print(count)


}
