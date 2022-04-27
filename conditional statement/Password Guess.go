// Write a program that reads a password (text) entered by the user and checks if the entered password matches the phrase "s3cr3t! P @ ssw0rd". 
// In case of coincidence, display "Welcome". In case of discrepancy, display "Wrong password!".

package main

import "fmt"

func main(){
	var password string
	fmt.Scanln(&password)

	if password == "s3cr3t!P@ssw0rd"{
		fmt.Println("Welcome")
	}else{
		fmt.Println("Wrong password!")
	}
}
