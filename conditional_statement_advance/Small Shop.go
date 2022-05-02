// An enterprising Bulgarian opens neighborhood shops in several cities and sells at different prices:
// city / product coffee water beer sweets peanuts
// Sofia           0.50  0.80  1.20  1.45   1.60
// Plovdiv         0.40  0.70  1.15  1.30   1.50
// Varna           0.45  0.70  1.10  1.35   1.55
// Write a program that reads product (string), city (string) and quantity (decimal number) entered by 
// the user and calculates and prints how much the corresponding quantity of the selected product costs in the specified city.

package main

import "fmt"

func main(){
	var product string
	fmt.Scanln(&product)

	var city string
	fmt.Scanln(&city)

	var quantity float32
	fmt.Scanln(&quantity)

	if city == "Sofia"{
		if product == "coffee"{
			fmt.Println(quantity * 0.50)
		}else if product == "water"{
			fmt.Println(quantity * 0.80)
		}else if product == "beer"{
			fmt.Println(quantity * 1.20)
		}else if product == "sweets"{
			fmt.Println(quantity * 1.45)
		}else if product == "peanuts"{
			fmt.Println(quantity * 1.60)
		}
	}else if city == "Plovdiv"{
		if product == "coffee"{
			fmt.Println(quantity * 0.40)
		}else if product == "water"{
			fmt.Println(quantity * 0.70)
		}else if product == "beer"{
			fmt.Println(quantity * 1.15)
		}else if product == "sweets"{
			fmt.Println(quantity * 1.30)
		}else if product == "peanuts"{
			fmt.Println(quantity * 1.50)
		}
	}else if city == "Varna"{
		if product == "coffee"{
			fmt.Println(quantity * 0.45)
		}else if product == "water"{
			fmt.Println(quantity * 0.70)
		}else if product == "beer"{
			fmt.Println(quantity * 1.10)
		}else if product == "sweets"{
			fmt.Println(quantity * 1.35)
		}else if product == "peanuts"{
			fmt.Println(quantity * 1.55)
		}
	}	
}
