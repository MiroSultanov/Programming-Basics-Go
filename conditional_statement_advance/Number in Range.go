// Да се напише програма, която проверява дали въведеното от потребителя число е в интервала [-100, 100] и е различно от 0 и извежда "Yes", 
// ако отговаря на условията, или "No" ако е извън тях.

package main

import "fmt"

func main(){
	var number int
	fmt.Scanln(&number)

	if number >= -100 && number <= 100 && number != 0{
		fmt.Print("Yes")
	}else{
		fmt.Print("No")
	}
}
