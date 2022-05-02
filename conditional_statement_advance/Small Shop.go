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