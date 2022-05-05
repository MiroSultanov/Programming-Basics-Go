// Climbers from all over Bulgaria gather in groups and mark the next peaks to climb. Depending on the size of the group, the climbers will climb different peaks.
// • Group of up to 5 people - climb Musala
// • Group of 6 to 12 people - climb Mont Blanc
// • Group of 13 to 25 people - climb Kilimanjaro
// • Group of 26 to 40 people - climb K2
// • A group of 41 or more people - climb Everest
// Write a program that calculates the percentage of climbers climbing each peak.

package main

import "fmt"

func main() {
	var groups int
	fmt.Scanln(&groups)

	var groupMusala int = 0
	var groupMonblan int = 0
	var groupKilimanjaro int = 0
	var groupK2 int = 0
	var groupEverest int = 0

	for group := 1; group <= groups; group++ {
		var countPeople int
		fmt.Scanln(&countPeople)
		if countPeople <= 5 {
			groupMusala += countPeople
		} else if countPeople <= 12 {
			groupMonblan += countPeople
		} else if countPeople <= 25 {
			groupKilimanjaro += countPeople
		} else if countPeople <= 40 {
			groupK2 += countPeople
		} else {
			groupEverest += countPeople
		}
	}
	var all int = groupEverest + groupK2 + groupKilimanjaro + groupMusala + groupMonblan
	var percentMusala float32 = float32(groupMusala) * 100 / float32(all)
	var percentMonblan float32 = float32(groupMonblan) * 100 / float32(all)
	var percentKilimanjaro float32 = float32(groupKilimanjaro) * 100 / float32(all)
	var percentK2 float32 = float32(groupK2) * 100 / float32(all)
	var percentEverest float32 = float32(groupEverest) * 100 / float32(all)
	
	fmt.Printf("%.2f%%\n", percentMusala)
	fmt.Printf("%.2f%%\n", percentMonblan)
	fmt.Printf("%.2f%%\n", percentKilimanjaro)
	fmt.Printf("%.2f%%\n", percentK2)
	fmt.Printf("%.2f%%", percentEverest)
}
