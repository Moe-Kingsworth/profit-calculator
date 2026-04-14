package main

import "fmt"

func CalculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 + taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func main() {

	revenue := getUserInput("Enter Revenue: ")
	expenses := getUserInput("Enter Expenses: ")
	taxRate := getUserInput("Enter Tax Rate: ")

	ebt, profit, ratio := CalculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.0f\n", ebt)
	fmt.Printf("Profit: %.0f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n.", ratio)
}
