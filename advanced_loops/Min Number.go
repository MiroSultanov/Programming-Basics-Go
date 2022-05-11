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