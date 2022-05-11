// # Write a program in which Marin solves exam tasks until he receives an "Enough" message from his lecturer. 
// # For each task solved, he receives a grade. The program must end reading data with the "Enough" command or if Marin receives the specified number of unsatisfactory
// # grades. Any grade less than or equal to 4 is unsatisfactory.

// # Input
// # ⦁ In the first row - number of unsatisfactory grades - integer;
// # ⦁ Then two lines are read repeatedly:
// # ⦁ Task name - text;
// # ⦁ Score - an integer in the range [2… 6]

// # Output
// # ⦁ If Marin reaches the "Enough" command, print in 3 lines:
// # ⦁ "Average score:"
// # ⦁ "Number of problems: {number of all tasks}"
// # ⦁ "Last problem: {name of last task}"
// # ⦁ If he receives the specified number of unsatisfactory grades:
// # } "You need a break, {number of unsatisfactory grades} poor grades."
// # The average score should be formatted to the second decimal place.

package main

import "fmt"

func main() {
	var badGrades int
	fmt.Scanln(&badGrades)

	var countBadGrades int = 0
	var countGrades int = 0
	var sumOfGrades int = 0
	var exit bool = false
	var nameOfTask string
	fmt.Scanln(&nameOfTask)
	var lastProblem string = ""

	for nameOfTask != "Enough" {
		var grade int 
		fmt.Scanln(&grade)

		if grade <= 4 {
			countBadGrades ++
		}

		if countBadGrades == badGrades {
			fmt.Printf("You need a break, %d poor grades.", badGrades)
			exit = true
			break
		}
		countGrades ++ 
		sumOfGrades += grade
		lastProblem = nameOfTask
		fmt.Scanln(&nameOfTask)
	}

	if !exit {
		var averageGrade float32 = float32(sumOfGrades) / float32(countGrades)
		fmt.Printf("Average score: %.2f\n",averageGrade)
		fmt.Printf("Number of problems: %d\n",countGrades)
		fmt.Printf("Last problem: %s",lastProblem)
	}
}
