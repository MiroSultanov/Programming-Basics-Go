// # Write a program that, before receiving the "Stop" command, reads integers entered by the user, finds the largest of them and prints it. Enter one number per line.

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var max int = math.MinInt

	var input string
	fmt.Scanln(&input)

	for input != "Stop" {
		number,_ := strconv.Atoi(input)

		if number > max {
			max = number
		}

		fmt.Scanln(&input)
	}

	fmt.Print(max)
}
