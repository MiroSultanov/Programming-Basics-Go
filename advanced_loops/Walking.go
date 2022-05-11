// # Gabby wants to start a healthy lifestyle and has set a goal of walking 10,000 steps every day. 
// # However, some days she is very tired from work and will want to go home before she achieves her goal. 
// # Write a program that reads from the console how many steps it takes each time it goes out during the day and when it achieves its goal of writing
// # "Goal reached! Good job!" and how many more steps she took "{the difference between the steps} steps over the goal!"
// # If she wants to go home before that, she will enter the "Going home" command and enter the steps she took while returning home. 
// # Then, if it fails to achieve its goal, the console should read: "{step difference} more steps to reach goal."

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
