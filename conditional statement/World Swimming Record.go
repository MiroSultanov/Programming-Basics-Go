package main

import "fmt"
import "math"

func main(){
	var recordInSeconds float32
	fmt.Scanln(&recordInSeconds)

	var distanceInMeters float32
	fmt.Scanln(&distanceInMeters)

	var timeInSecondForMeter float32
	fmt.Scanln(&timeInSecondForMeter)

	var needToSwim float32 = distanceInMeters * timeInSecondForMeter
	var delay = math.Floor(float64(distanceInMeters)/ 15.0) * 12.5
	var totalTime float32 = needToSwim + float32(delay)

	if totalTime < recordInSeconds{
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", totalTime)
	}else{
		var neededTime float32 = totalTime - recordInSeconds
		fmt.Printf("No, he failed! He was %.2f seconds slower.",neededTime )
	}


}