// # Write a program that reads an integer N entered by the user and generates all possible "special" numbers from 1111 to 9999. 
// # To be a "special" number, it must meet the following condition:
// # ⦁ N to be divisible by each of its digits without remainder.
// # Example: at N = 16, 2418 is a special number:
// # ⦁ 16/2 = 8 without residue
// # ⦁ 16/4 = 4 without residue
// # ⦁ 16/1 = 16 without residue
// # ⦁ 16/8 = 2 without residue

// # Input
// # The input is read from the console and consists of an integer in the range [1… 600000]

// # Output
// # All "special" numbers must be printed on the console, separated by a space

package main

import "fmt"

func main2() {

	var n int = 9
	

	for thousands := 1; thousands <= 9; thousands++ {
		for hundreds := 1; hundreds <= 9; hundreds++ { 
			for tens := 1; tens <= 9; tens++ {       
				for units := 1; units <= 9; units++ { 
					
					var check1 bool = n % units == 0
					var check2 bool = n % tens == 0
					var check3 bool = n % hundreds == 0
					var check4 bool = n % thousands == 0
					if check1 && check2 && check3 && check4 {
						fmt.Printf("%d%d%d%d ", thousands, hundreds, tens, units)
					}
				}
			}
		}
	}
}
