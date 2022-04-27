// Write a program that reads the speed (real number) entered by the user and prints speed information.
// • At speeds up to 10 (inclusive) print "slow"
// • At speeds above 10 and up to 50 (inclusive) print "average"
// • At speeds above 50 and up to 150 (inclusive) print "fast"
// • At speeds above 150 and up to 1000 (inclusive) print "ultra fast"
// • Print "extremely fast" at higher speeds

package main

import "fmt"

func main(){
	var speed float32
fmt.Scanln(&speed)

if speed <= 10 {
	fmt.Println("slow")
}else if speed <= 50 {
	fmt.Println("average")
}else if speed <= 150 {
	fmt.Println("fast")
}else if speed <= 1000 {
	fmt.Println("ultra fast")
}else{
	fmt.Println("extremely fast")
}
}
