package main

import (
	"fmt"
	"strings"
)

func Exercise1_Extremes() {
	var count int

	fmt.Print("Enter the number of integers you want to input: ")
	fmt.Scanf("%d", &count)

	if count <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	minimum := int(^uint(0) >> 1)
	maximum := int(^uint(0)>>1) - 1

	for counter := 0; counter < count; counter++ {
		var number int
		fmt.Print("Enter number ", counter+1, ": ")
		fmt.Scanf("%d", &number)

		if counter == 0 {
			minimum = number
			maximum = number
		} else {
			if number < minimum {
				minimum = number
			}
			if number > maximum {
				maximum = number
			}
		}
	}

	fmt.Println("Minimum:", minimum)
	fmt.Println("Maximum:", maximum)

}

func Exercise2_integerDivisibleBy3() {
	sum := 0
	for count := 1; count <= 30; count++ {
		if count%3 == 0 {
			sum += count
		}
	}
	fmt.Println("Sum of integers divisible by 3 from 1 to 30:", sum)

}

func Exercise3_sumOfASeries() {
	var sum int64

	fmt.Printf("%-10s %s\n", "n", "Sum")
	fmt.Println("---------------------")
	for count := int64(1); count <= 100; count++ {
		sum += count
		fmt.Printf("%-10d %d\n", count, sum)

	}
}

func Exercise4_trianglePrinting() {
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

func Exercise5_studentGrade() {
	var grade string
	var name string
	var aCount, bCount, cCount, dCount int

	for i := 1; i <= 5; i++ {
		fmt.Printf("Enter name of student %d: ", i)
		fmt.Scanln(&name)

		fmt.Printf("Enter grade for %s (A, B, C, D): ", name)
		fmt.Scanln(&grade)

		grade = strings.ToUpper(grade)

		switch grade {
		case "A":
			aCount++
		case "B":
			bCount++
		case "C":
			cCount++
		case "D":
			dCount++
		default:
			fmt.Println("Invalid grade entered. Only A, B, C, or D allowed.")

		}
	}

	fmt.Println("\nGrade Summary:")
	fmt.Println("Number of A's:", aCount)
	fmt.Println("Number of B's:", bCount)
	fmt.Println("Number of C's:", cCount)
	fmt.Println("Number of D's:", dCount)
}

func Exercise6_diamondPrintingProgram() {
	n := 5

	for count := 1; count <= n; count++ {

		for innerCounter := 1; innerCounter <= n-count; innerCounter++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*count-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for count := n - 1; count >= 1; count-- {

		for innerCounter := 1; innerCounter <= n-count; innerCounter++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*count-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	// Exercise1_Extremes()
	// Exercise2_integerDivisibleBy3()
	// Exercise3_sumOfASeries()
	// Exercise4_trianglePrinting()
	// Exercise5_studentGrade()
	Exercise6_diamondPrintingProgram()

}
