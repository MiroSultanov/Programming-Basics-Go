// Write a program that reads two integers (N1 and N2) and an operator with which to perform a mathematical operation with them. Possible operations are: 
// Addition (+), Subtraction (-), Multiplication (*), Division (/) and Division with remainder (%). When adding, subtracting and multiplying the console, 
// the result must be printed and whether it is even or odd. In ordinary division - the result. In modular division - the remainder. It should be borne in mind that 
// the divisor can be equal to 0 (zero) and not divisible by zero. In this case, a special message must be printed.

package main

import "fmt"

func main () {

	var n1, n2 int
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)

	var operator string
	fmt.Scanln(&operator)


	switch operator {
		case "+":
			var sum int = n1 + n2
			if sum % 2 == 0 {

				fmt.Printf("%d + %d = %d - even", n1, n2, sum)
			} else {

				fmt.Printf("%d + %d = %d - odd", n1, n2, sum)
			}
		case "-":
			var diff int = n1 - n2
			if diff % 2 == 0 {
				fmt.Printf("%d - %d = %d - even", n1, n2, diff)
			} else {
				fmt.Printf("%d - %d = %d - odd", n1, n2, diff)
			}
		case "*":
			var product int = n1 * n2
			if product % 2 == 0 {
				fmt.Printf("%d * %d = %d - even", n1, n2, product)
			} else {
				fmt.Printf("%d * %d = %d - odd", n1, n2, product)
			}
		case "/":
			if n2 == 0 {
				fmt.Printf("Cannot divide %d by zero", n1)
			} else {
				var division float32 = float32(n1) / float32(n2)
				fmt.Printf("%d / %d = %.2f", n1, n2, division)
			}
	
		case "%":
			if n2 == 0 {
				fmt.Printf("Cannot divide %d by zero", n1)
			} else {
				var reminder int = n1 % n2
				fmt.Printf("%d %% %d = %d", n1, n2, reminder)
			}
	}

}
