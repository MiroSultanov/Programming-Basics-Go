package main

import (
	"fmt"
	"strconv"
)

func main() {
	var targetSteps int = 10000

	var totalWalkedSteps int = 0

	var input string
	fmt.Scanln(&input)

	for input != "GoingHome" {
		walkingSteps, _ := strconv.Atoi(input)

		totalWalkedSteps += walkingSteps

		if totalWalkedSteps >= targetSteps {
			fmt.Println("Goal reached! Good job!")
			fmt.Printf("%d steps over the goal!", totalWalkedSteps - targetSteps)
		}

		fmt.Scanln(&input)
	}
	
	if input == "GoingHome" {
		var stepsToHome int
		fmt.Scanln(&stepsToHome)
		totalWalkedSteps += stepsToHome

		if totalWalkedSteps >= targetSteps {
			fmt.Println("Goal reached! Good job!")
			fmt.Printf("%d steps over the goal!", totalWalkedSteps - targetSteps)
		}else {
			fmt.Printf("%d steps over the goal!", totalWalkedSteps - targetSteps)
		}
	}
}