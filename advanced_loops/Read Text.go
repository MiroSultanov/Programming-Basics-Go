// # Write a program that reads text from the console (string) and prints it until it receives the "Stop" command.

package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	for input != "Stop" {
		fmt.Println(input)
		fmt.Scanln(&input)
	}
}
