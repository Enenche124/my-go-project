package main

import "fmt"

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
	for count := int64(1); count <= 100; count++ {
		sum += count
		fmt.Printf("%-10d %d\n", count, sum)

	}
}
func main() {
	Exercise1_Extremes()
	Exercise2_integerDivisibleBy3()
	Exercise3_sumOfASeries()
}
