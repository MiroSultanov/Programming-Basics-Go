package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	var area int = a * b
	fmt.Println(area)
}
