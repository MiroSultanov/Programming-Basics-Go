// An integer is given - the starting number of points. Bonus points are accrued on it according to the rules described below. 
// Write a program that calculates the bonus points received by the number and the total number of points (number + bonus).
// • If the number is up to 100 inclusive, the bonus points are 5.
// • If the number is greater than 100, the bonus points are 20% of the number.
// • If the number is greater than 1000, the bonus points are 10% of the number.

// • Additional bonus points (accrued separately from the previous ones):
// o For an even number  + 1 point.
// o For a number ending in 5  + 2 points.

package main

import "fmt"

func main(){
	var startPoint int
	fmt.Scanln(&startPoint)
	var bonus float32 = 0 

	if startPoint <= 100{
		bonus = 5
	}else if startPoint > 1000{
		bonus = float32(startPoint) * 0.1
	}else {
		bonus = float32(startPoint) * 0.2
	}
	
	if startPoint % 2 == 0{
		bonus = bonus + 1
	}else if startPoint % 10 == 5{
		bonus = bonus + 2
	}

	fmt.Println(bonus)
	fmt.Println(float32(startPoint) + bonus)
}
