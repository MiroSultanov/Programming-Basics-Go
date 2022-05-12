// # Your task is to write a program that calculates the percentage of tickets for each type of ticket sold: student (student), standard (standard) and child (kid),
// # for all screenings. You also need to calculate what percentage of the hall is filled for each screening.

// # Input
// # The input is a series of integers and text:
// # ⦁ On the first line until you receive the command "Finish" - movie name - text
// # ⦁ In the second row - the free seats in the hall for each screening - integer [1… 100]
// # ⦁ For each film, one line is read until the vacancies in the hall are exhausted or until the "End" command is received:
// # ⦁ Type of ticket purchased - text ("student", "standard", "kid")

// # Output
// # The following lines must be printed on the console:
// # ⦁ After each film, print what percentage of the cinema is full
// # "{movie name} - {percentage of hall full}% full."
// # ⦁ When receiving the "Finish" command, print four lines:
// # ⦁ "Total tickets: {total number of tickets purchased for all films}"
// # } "{Percentage of student tickets}% student tickets."
// # ⦁ "{percentage of standard tickets}% standard tickets."
// # } "{Percentage of children's tickets}% kids tickets."

package main

import "fmt"

func main() {
	var countStudent int = 0
	var countStandard int = 0
	var countKid int = 0
	for true {
		var nameTillFinish string
		fmt.Scanln(&nameTillFinish)
		if nameTillFinish == "Finish" {
			break
		}
		var countAll int = 0
		var places int
		fmt.Scanln(&places)
		for true {
			var typeTillEnd string
			fmt.Scanln(&typeTillEnd)
			if typeTillEnd == "End" {
				break
			}
			switch typeTillEnd {
			case "student":
				countStudent++
				break
			case "standard":
				countStandard++
				break
			case "kid":
				countKid++
				break
			}
			countAll++
			if countAll >= places {
				break
			}

		}
		var percentFull float32 = float32(countAll) * 100 / float32(places)
		fmt.Printf("%s - %.2f%% full.\n", nameTillFinish, percentFull)
		countAll = 0
	}
	var total int = countKid + countStandard + countStudent
	var percentStudent float32 = float32(countStudent) * 100 / float32(total)
	var percentStandard float32 = float32(countStandard) * 100 / float32(total)
	var percentKid float32 = float32(countKid) * 100 / float32(total)
	fmt.Printf("Total tickets: %d\n", total)
	fmt.Printf("%.2f%% student tickets.\n", percentStudent)
	fmt.Printf("%.2f%% standard tickets.\n", percentStandard)
	fmt.Printf("%.2f%% kids tickets.", percentKid)
}
