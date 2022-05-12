// # The "Train the trainers" course is coming to an end and the final evaluation is approaching. Your task is to help the jury that will 
// # evaluate the presentations by writing a program in which to calculate the average grade from the performance of each presentation by a student, 
// # and finally - the average success of all of them.
// # From the console in the first row is read the number of people on the jury n - an integer.
// # Then the name of the presentation is read on a separate line - text.
// # For each presentation on a new line read n - the number of marks - a real number.
// # After calculating the average score for a specific presentation, the following is printed on the console:
// #  "{presentation name} - {average grade}."
// # After receiving the "Finish" command, the console reads "Student's final assessment is {average of all presentations}." and the program ends.
// # All grades must be formatted to the second decimal place.

package main

import "fmt"

func main0 () {

	var countJury int 
	fmt.Scanln(&countJury)


	var input string
	fmt.Scanln(&input)

	var sumAllGrades float32 = 0 
	var countAllGrades int = 0 

	for input != "Finish" {
		
		var presentation string = input
		var sumGradesPerPresentation float32 = 0 
		
		for jury := 1; jury <= countJury; jury++ {
			var grade float32
			fmt.Scanln(&grade)
			
			sumGradesPerPresentation += grade
			
			sumAllGrades += grade
			countAllGrades++
		}
		
		var average float32 = sumGradesPerPresentation / float32(countJury)
		fmt.Printf("%s - %.2f.\n", presentation, average)

		fmt.Scanln(&input) 
	}

	
	var averageAll float32 = sumAllGrades / float32(countAllGrades)
	fmt.Printf("Student's final assessment is %.2f.", averageAll)




}
