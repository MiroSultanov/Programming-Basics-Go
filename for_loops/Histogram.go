// N integers are given in the interval [1… 1000]. Of these, some percentage p1 are below 200, another percentage p2 are from 200 to 399, another percentage 
// p3 are from 400 to 599, another percentage p4 are from 600 to 799 and the remaining p5 percent are from 800 up. 
// Write a program that calculates and prints the percentages p1, p2, p3, p4 and p5.

package main

import "fmt"

func main(){
	var n int 
	fmt.Scanln(&n)

	var group1, group2, group3, group4, group5 int = 0, 0, 0, 0, 0

	for number :=1; number <= n; number ++ {
		var value int 
		fmt.Scanln(&value)

		if value < 200 {
			group1++
		}else if value >= 200 && value <= 399 {
			group2++
		}else if value >= 400 && value <= 599 {
			group3++
		}else if value >= 600 && value <= 799 {
			group4++
		}else if value >= 800 {
			group5++
		}
	}

	var p1 float32 = float32(group1) / float32(n) * 100.0
	var p2 float32 = float32(group2) / float32(n) * 100.0
	var p3 float32 = float32(group3) / float32(n) * 100.0
 	var p4 float32 = float32(group4) / float32(n) * 100.0
	var p5 float32 = float32(group5) / float32(n) * 100.0

	fmt.Printf("%.2f%%\n",p1)
	fmt.Printf("%.2f%%\n",p2)
	fmt.Printf("%.2f%%\n",p3)
	fmt.Printf("%.2f%%\n",p4)
	fmt.Printf("%.2f%%\n",p5)
}
