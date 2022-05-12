// # Write a program that prints the hours of the day from 0: 0 to 23:59, each on a separate line. The hours must be displayed in the format "{hour}: {minutes}".

package main

import "fmt"

func main() {

	for hour := 0; hour <= 23; hour++ {
		for minutes := 0; minutes <= 59; minutes++ {
			fmt.Printf("%d:%d\n", hour, minutes)
		}
	}
}
