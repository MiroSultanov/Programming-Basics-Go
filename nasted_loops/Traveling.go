// # Annie loves to travel and wants to visit several different destinations this year. When she chooses a destination, she will decide how much money 
// # she will need to get there, and she will start saving. When she has saved enough, she will be able to travel.
// # The destination and the minimum budget (decimal number) that will be needed for the trip will be read from the console each time.
// # Then a few sums (decimal numbers) will be read, which Annie saves by working and when she manages to collect enough for the trip,
// # she will leave, and the console should say: "Going to {destination}!"
// # When she has visited all the destinations she wants, instead of a destination she will enter "End" and the program will end. 

package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	for input != "End" {
		var needMoney float32
		fmt.Scanln(&needMoney)

		var savedMoney float32 = 0

		for savedMoney < needMoney {
			var money float32
			fmt.Scanln(&money)

			savedMoney += money
		}
		fmt.Printf("Going to %s!\n", input)

		fmt.Scanln(&input)
	}
}
