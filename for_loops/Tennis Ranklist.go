// Grigor Dimitrov is a tennis player whose next goal is to climb the world tennis rankings for men.
// During the year, Grisho participates in a certain number of tournaments, and for each tournament receives points that depend on the position on which he ended 
// in the tournament. There are three options for completing a tournament:
//  W - If a winner gets 2000 points
//  F - If a finalist gets 1200 points
//  SF - If it is semi -finalist receives 720 points
// Write a program that calculates how many Grigor points will be after playing all tournaments, knowing how many points starts the seasons.
// Also calculate how many points win from all tournaments played and how many percent of tournaments he has won.

package main

import "fmt"

func main() {
	var tournaments, startPoints int
	fmt.Scanln(&tournaments)
	fmt.Scanln(&startPoints)

	var points int = 0
	var wins int = 0
	for i := 1; i <= tournaments; i++ {
		var place string
		fmt.Scanln(&place)
		switch place {
		case "W":
			points += 2000
			wins++
			break
		case "F":
			points += 1200
			break
		case "SF":
			points += 720
			break
		}
	}
	var allPoints int = startPoints + points
	var averagePoints int = points / tournaments
	var percentWins float32 = float32(wins) * 100 / float32(tournaments)

	fmt.Printf("Final points: %d\n", allPoints)
	fmt.Printf("Average points: %d\n", averagePoints)
	fmt.Printf("%.2f%%", percentWins)
}
