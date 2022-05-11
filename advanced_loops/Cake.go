// # You are invited to the 30th birthday, when the birthday boy eats a huge cake. However, he does not know how many pieces his guests can take from her. 
// # Your task is to write a program that calculates the number of pieces that guests have taken before it ends. You will get the dimensions of the 
// # cake (width and length - integers and then on each line, until you receive the command "STOP" or until the cake is finished, the number of 
// # pieces that guests take from it. Each piece of cake is 1x1 cm .
// # Print one of the following lines on the console:
// # "{Number of pieces} pieces are left." - if you reach STOP and the pieces of cake are not finished;
// #  "No more cake left! You need {number of missing pieces} pieces more."

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main1 () {

	var width, lenght int 
	fmt.Scanln(&width) 
	fmt.Scanln(&lenght)
	
	var pieces int = width * lenght


	var input string
	fmt.Scanln(&input)

	for input != "STOP" {
		
		takenPieces, _ := strconv.Atoi(input) //колко парчета са взети

		
		pieces -= takenPieces

		
		if pieces <= 0 {
			
			fmt.Printf("No more cake left! You need %.0f pieces more.", math.Abs(float64(pieces)))
			break
		} 

		fmt.Scanln(&input)

	}

	if input == "STOP" {
		fmt.Printf("%d pieces are left.", pieces)
	}
}
