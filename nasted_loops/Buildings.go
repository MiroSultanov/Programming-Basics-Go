// # Write a program that displays the numbers of the rooms in a building (in descending order) on the console, provided that the following conditions are met:
// # ⦁ There are only offices on each even floor;
// # ⦁ There are only apartments on each odd floor;
// # ⦁ Each apartment is denoted as follows: A {floor number} {apartment number}, apartment numbers start from 0;
// # ⦁ Each office is denoted as follows: О {floor number} {office number}, office numbers also start from 0;
// # ⦁ There are always apartments on the top floor and they are bigger than the others, so their number says 'L' instead of 'A'. 
// # If there is only one floor, there are only large apartments!
// # Two integers are read from the console - the number of floors and the number of rooms per floor.

package main

import "fmt"

func main() {
	var floors int 
	fmt.Scanln(&floors)

	var rooms int
	fmt.Scanln(&rooms)

	for floor := floors; floor >= 1; floor--{
		for room := 0; room <= rooms - 1; room++ {
			if floor == floors {
				fmt.Printf("L%d%d ", floor, room)
			}else if floor % 2 == 0 {
				fmt.Printf("O%d%d ", floor, room)
			}else if floor % 2 != 0 {
				fmt.Printf("A%d%d ",floor ,room)
			}
		}
		fmt.Println()
	}
}
