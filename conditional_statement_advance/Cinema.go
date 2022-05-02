// In a movie theater, the chairs are arranged in a rectangular shape in r rows and c columns. There are three types of screenings with tickets at different prices:
// • Premiere - premiere screening, at a price of BGN 12.00.
// • Normal - standard screening, at a price of BGN 7.50.
// • Discount - screening for children, pupils and students at a reduced price of BGN 5.00.
// Write a program that reads the type of projection (string), number of rows and number of columns in the hall (integers) entered by the user, 
// and calculates the total ticket revenue for a full hall. Print the result in the format as in the examples below, 2 characters after the decimal point.

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
