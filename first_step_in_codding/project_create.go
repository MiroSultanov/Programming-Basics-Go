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
