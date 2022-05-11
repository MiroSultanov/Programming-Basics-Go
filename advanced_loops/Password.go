// # Write a program that initially reads a username and password. Then read the login password.
// # ⦁ when entering the wrong password: prompt the user to enter a new password.
// # ⦁ when entering the correct password: print "Welcome {username}!".

package main

import "fmt"

func main() {
	var username string 
	fmt.Scanln(&username)

	var password string
	fmt.Scanln(&password)

	var enterPass string
	fmt.Scanln(&enterPass)

	for enterPass != password {
		fmt.Scanln(&enterPass)
	}

	fmt.Printf("Welcome %s!", username)

}
