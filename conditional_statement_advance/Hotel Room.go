// The hotel offers 2 types of rooms: studio and apartment. Write a program that calculates the price for the entire stay for a studio and apartment. 
// Prices depend on the month of stay:
// May and October                  June and September               July and August
// Studio - BGN 50 / night        Studio - BGN 75.20 / night      Studio - BGN 76 / night
// Apartment - 65 BGN / night     Apartment - 68.70 BGN / night   Apartment - 77 BGN / night
// The following discounts are also available:
// • For studio, for more than 7 nights in May and October: 5% discount.
// • For studio, for more than 14 nights in May and October: 30% discount.
// • For studio, for more than 14 nights in June and September: 20% discount.
// • For an apartment, for more than 14 nights, regardless of the month: 10% discount.

package main

import "fmt"

func main() {
	var month string
	fmt.Scanln(&month)

	var count int
	fmt.Scanln(&count)

	var totalStudio float32 = 0
	var totalApartment float32 = 0

	switch month {
		case "May", "October":
			totalStudio = float32(count) * 50
			totalApartment = float32(count) * 65

			if count > 7 && count < 15 {
				totalStudio = totalStudio * 0.95
			} else if count > 14 {
				totalStudio = totalStudio * 0.7
				totalApartment = totalApartment * 0.9
			}
			break
		case "June", "September":
			totalStudio = float32(count) * 75.2
			totalApartment = float32(count) * 68.7
			if count > 14 {
				totalStudio = totalStudio * 0.8
				totalApartment = totalApartment * 0.9
			}
			break
		case "July", "August":
			totalStudio = float32(count) * 76
			totalApartment = float32(count) * 77
			if count > 14 {
				totalApartment = totalApartment * 0.9
			}
	}
	fmt.Printf("Apartment: %.2f lv.\n", totalApartment)
	fmt.Printf("Studio: %.2f lv.", totalStudio)
}
