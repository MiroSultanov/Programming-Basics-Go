// # Ani goes to her hometown after a very long period abroad. On her way home, she sees her grandmother's old library and remembers her favorite book. 
// # Help Annie by writing a program in which she enters the book (text) she is looking for. Until Annie finds her favorite book or checks all the books in the library, 
// # the program must read each time on a new line the name of each subsequent book (text) that she checks. The books in the library are gone as soon as you 
// receive the text "No More Books".
// # ⦁ If he does not find the book he is looking for, print it in two lines:
// # ⦁ "The book you search is not here!"
// # ⦁ "You checked {number} books."
// # ⦁ If he finds his book, one line is printed:
// # ⦁ "You checked {number} books and found it."

package main

import "fmt"

func main() {
	var serachBook string
	fmt.Scanln(&serachBook)

	var inputBook string 
	fmt.Scanln(&inputBook)

	var count int = 0
	var isFound bool = false

	for inputBook != "NoMoreBooks" {
		if inputBook == serachBook {
			fmt.Printf("You checked %d books and found it.", count)
			isFound =  true
			break
		}
		count ++
		fmt.Scanln(&inputBook)
	}

	if !isFound {
		fmt.Println("The book you search is not here!")
		fmt.Printf("You checked %d books.", count)
	}

}
