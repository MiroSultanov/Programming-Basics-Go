package main

import "fmt"

func main(){
	var typeProjction string
	fmt.Scanln(&typeProjction)

	var rows int 
	fmt.Scanln(&rows)

	var column int 
	fmt.Scanln(&column)

	var countTicket int = rows * column

	var pricePerTicket float32 = 0

	switch typeProjction{
	case "Premiere":
		pricePerTicket = 12.00
	case "Normal":
		pricePerTicket = 7.50
	case "Discount":
		pricePerTicket = 5.00
	}

	var profit float32 = float32(countTicket) * pricePerTicket

	fmt.Printf("%.2f leva", profit)
}