// The company gives the following commissions to its merchants according to the city in which they work and the volume of sales s:

// town	0 ≤ s ≤ 500	500 < s ≤ 1 000  	1 000 < s ≤ 10 000	s > 10 000
// Sofia	    5%	               7%	               8%	            12%
// Varna	   4.5%         	7.5%	              10%	            13%
// Plovdiv	   5.5%	               8%	              12%	           14.5% 

// Write a console program that reads the city name (text) and sales volume (real number) entered by the user, and calculates and displays the amount of the 
// trade commission according to the table above. Display the result formatted to 2 digits after the decimal point. In case of invalid city or sales 
// volume (negative number) print "error".

package main

import "fmt"

func main(){
	var city string
	fmt.Scanln(&city)

	var sales float32
	fmt.Scanln(&sales)

	var commission float32

	if city == "Sofia"{
		if sales >= 0 && sales <= 500{
			commission = sales * 0.05
		}else if sales > 500 && sales <= 1000{
			commission = sales * 0.07
		}else if sales > 1000 && sales <= 10000{
			commission = sales * 0.08
		}else if sales > 10000{
			commission = sales * 0.12
		}
	}else if city == "Varna"{
		if sales >= 0 && sales <= 500{
			commission = sales * 0.045
		}else if sales > 500 && sales <= 1000{
			commission = sales * 0.075
		}else if sales > 1000 && sales <= 10000{
			commission = sales * 0.10
		}else if sales > 10000{
			commission = sales * 0.13
		}
	}else if city == "Plovdiv"{
		if sales >= 0 && sales <= 500{
			commission = sales * 0.055
		}else if sales > 500 && sales <= 1000{
			commission = sales * 0.08
		}else if sales > 1000 && sales <= 10000{
			commission = sales * 0.12
		}else if sales > 10000{
			commission = sales * 0.145
		}
	}
	
	if commission == 0{
		fmt.Println("error")
	}else{
		fmt.Printf("%.2f", commission)
	}
}
