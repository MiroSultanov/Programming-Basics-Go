// # Write a program that, before receiving the "Stop" command, reads integers entered by the user, finds at least one of them and prints it. Enter one number per line.

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var min int = math.MaxInt

	var input string
	fmt.Scanln(&input)

	for input != "Stop" {
		number,_ := strconv.Atoi(input)

		if number > min {
			min = number
		}

		fmt.Scanln(&input)
	}

	fmt.Print(min)
}
