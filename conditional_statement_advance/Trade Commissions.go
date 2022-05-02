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