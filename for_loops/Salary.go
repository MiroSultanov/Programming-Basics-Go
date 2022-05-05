// A company boss notices that more and more employees are spending time on sites that distract them.
// To prevent this, he introduces surprise checks on the open tabs of his employees' browsers.
// According to the open site, the following fines are imposed in the tab:
// • "Facebook" -> BGN 150.
// • "Instagram" -> BGN 100.
// • "Reddit" -> BGN 50

package main

import "fmt"

func main() {
	var n, salary int
	fmt.Scanln(&n)
	fmt.Scanln(&salary)

	for tab := 1; tab <= n; tab++ {
		var site string
		fmt.Scanln(&site)
		
		switch site {
		case "Facebook":
			salary -= 150
			break
		case "Instagram":
			salary -= 100
			break
		case "Reddit":
			salary -= 50
			break
		}
		if salary <= 0 {
			fmt.Println("You have lost your salary.")
			break
		}
	}
	if salary > 0 {
		fmt.Println(salary)
	}
}
