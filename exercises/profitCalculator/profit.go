package main

import (
	"fmt"
)

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print("Please enter the your revenue amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter the your expenses amount: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter the your tax rate percentage: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println("Your earnings before tax are:", earningsBeforeTax)
	fmt.Println("Your profit will be:", profit)
	fmt.Println("And the ratio between earnings and profit is:", ratio)

}
