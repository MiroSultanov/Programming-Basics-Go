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