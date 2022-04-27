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