// Write a program that calculates how many hours it will take an architect to design several construction projects. The preparation of a project takes three hours.

// Input
// 2 lines are read from the console:
// 1. The name of the architect - text
// 2. Number of projects to be prepared - integer in the range [0… 100]

// Output
// The following is printed on the console:
// • "The architect {name of architect} will need {required hours} hours to complete {number of projects} project / s."

package main

import "fmt"

func main() {
	var arhitectName string
	fmt.Scanln(&arhitectName)
	var numberProjects int
	fmt.Scanln(&numberProjects)
	var needTime int = numberProjects * 3

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.", arhitectName, needTime, numberProjects)
}
