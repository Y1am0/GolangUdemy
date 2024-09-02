package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Please enter the desired investement amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years you will be investing for: ")
	fmt.Scan(&years)

	fmt.Print("Finally, enter the expected return rate for your specific investment: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Adjusted value: %.1f\n", futureRealValue)

	// fmt.Println("The future value of the investment is:", futureValue)
	// fmt.Println("The future value of the investement, adjusted for inflation is:", futureRealValue)
	// fmt.Printf("The future value of the investment is: %.2f\nThe future value adjusted for inflation is: %.2f", futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}
