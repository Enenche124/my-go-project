package main

import (
	"fmt"
)

// func table() {
// 	fmt.Printf("%-5s%-5s%-5s%-5s\n", "N", "N2", "N3", "N4")

// 	for count := 1; count <= 5; count++ {
// 		n := count
// 		n2 := n * n
// 		n3 := n * n * n
// 		n4 := n * n * n * n
// 		fmt.Printf("%-5d%-5d%-5d%-5d\n", n, n2, n3, n4)
// 	}
// }

// func palindrome(value int) bool {
// 	original := value
// 	reversed := 0

// 	for value > 0 {
// 		remainder := value % 10
// 		reversed = reversed*10 + remainder
// 		value /= 10
// 	}

// 	return original == reversed
// }

func trianglePrinting() {
	const rows = 10

	for count := 1; count <= rows; count++ {
		for innerCounter := 1; innerCounter <= count; innerCounter++ {
			fmt.Print("*")
		}
		for innerCounter := count + 1; innerCounter <= rows; innerCounter++ {
			fmt.Print(" ")
		}
		fmt.Print("  ")

		for innerCounter := count; innerCounter <= rows; innerCounter++ {
			fmt.Print("*")
		}

		for innerCounter := 1; innerCounter < count; innerCounter++ {
			fmt.Print(" ")
		}
		fmt.Print("  ")

		for innerCounter := 1; innerCounter < count; innerCounter++ {
			fmt.Print(" ")
		}
		for innerCounter := count; innerCounter <= rows; innerCounter++ {
			fmt.Print("*")
		}
		fmt.Print("  ")

		for innerCounter := 1; innerCounter <= rows-count; innerCounter++ {
			fmt.Print(" ")
		}
		for innerCounter := 1; innerCounter <= count; innerCounter++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func main() {
	// largest := math.MinInt64
	// secondLargest := math.MinInt64
	// counter := 0

	// for count := range 10 {
	// 	fmt.Print("Enter number ", count+1, " : ")
	// 	var number int
	// 	fmt.Scanf("%d", &number)
	// 	if number > largest {
	// 		secondLargest = largest
	// 		largest = number
	// 	} else if number > secondLargest {
	// 		secondLargest = number
	// 	}

	// 	counter++
	// 	if counter == 10 {
	// 		break
	// 	}

	// }
	// fmt.Println("The largest number is:", largest)
	// fmt.Println("The second largest number is:", secondLargest)

	// table()
	trianglePrinting()
	// fmt.Println("Palindrome check for 12321,:", palindrome(12321))
	// fmt.Println("Palindrome check for 123:", palindrome(123))
}
