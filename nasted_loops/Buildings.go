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