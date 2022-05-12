// # Print on the console the multiplication table for the numbers 1 to 10 in the format:
// # "{first multiplier} * {second multiplier} = {result}".

package main

import "fmt"


func main() {

	for first := 1; first <= 10; first++ {
		
		for second := 1; second <= 10; second++ {
			fmt.Printf("%d * %d = %d\n", first, second, first * second)
		}
	}
}
