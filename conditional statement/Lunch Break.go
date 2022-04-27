// During the lunch break you want to watch an episode of your favorite series. Your task is to write a program that will find out if you have enough time to watch
// the episode. During the holiday you spend time for lunch and time for rest. Lunch time will be 1/8 of the rest time, and rest time will be 1/4 of the rest time.

// Input
// 3 lines are read from the console:
// 1. Name of the series - text
// 2. Episode duration - integer in the range [10… 90]
// 3. Duration of the break - an integer in the range [10… 120]

// Output
// Write one line on the console:
// • If you have enough time to watch the episode:
// "You have enough time to watch {series name} and left with {free time} minutes free time."
// • If you do not have enough time:
// "You don't have enough time to watch {series name}, you need {time} more minutes."
// Time to round to the nearest whole number up.

package main

import "fmt"
import "math"

func main(){
	var seriesName string
	fmt.Scanln(&seriesName)

	var episodeDuration int
	fmt.Scanln(&episodeDuration)

	var restDuration int
	fmt.Scanln(&restDuration)

	var timeForLunch float32 = float32(restDuration) / 8.0
	var timeForRest float32 = float32(restDuration) / 4.0
	var leftTime float32 = float32(restDuration) - timeForLunch - timeForRest

	if leftTime >= float32(episodeDuration){
		var coolTime float32 = leftTime - float32(episodeDuration)
		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.", seriesName, math.Ceil(float64(coolTime)))
	}else{
		var neededTime float32 = float32(episodeDuration) - leftTime
		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.", seriesName, math.Ceil(float64(neededTime)))
	}
}
