// Write a program that reads text (person's name) from the console and prints "Hello, <name>!", Where <name> is the name entered from the console.

package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + "!")
}
