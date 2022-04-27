// Ivan decides to improve the World Record for long distance swimming. The console introduces the record in seconds that Ivan has to improve, the distance in meters
// he has to swim and the time in seconds for which he swims a distance of 1 m. To write a program that calculates whether he has done the task, such as it is considered
// that: the resistance of the water slows it down every 15 m by 12.5 seconds. When calculating how many times Ivan will slow down as a result of water resistance, 
// the result should be rounded down to the nearest whole number. To calculate the time in seconds for which Ivan will swim the distance and the difference from 
// the World Record.

// Input
// 3 lines are read from the console:
// 1. The record in seconds - a real number in the interval [0.00… 100000.00]
// 2. The distance in meters - a real number in the interval [0.00… 100000.00]
// 3. The time in seconds for which a distance of 1 m floats - a real number in the interval [0.00… 1000.00]

// Output
// Printing the console depends on the result:
// • If Ivan has improved the World Record (his time is less than the record) we print:
// o "Yes, he succeeded! The new world record is {Ivan's time} seconds."
// • If he has NOT improved the record (his time is greater than or equal to the record) we print:
// o "No, he failed! He was {seconds not slow} seconds slower."
// The result must be formatted to the second decimal place.

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
