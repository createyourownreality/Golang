package main

import "fmt"

func main() {

	Revenue := getUserInput("Enter the Revenue: ")
	Expenses := getUserInput("Enter the Expenses: ")
	TaxRate := getUserInput("Enter the TaxRate: ")

	EBT, Profit, Ratio := calculatevalues(Revenue, Expenses, TaxRate)

	RatioPrint := fmt.Sprintf("The Ratio of the (EBT/Profit) is: %0.2f", Ratio)

	// Using backticks we can print the multiline Strings
	//Printing all the results
	fmt.Printf("The EBT is %v: \n", EBT) // here we are using the Printf and parameters
	fmt.Println("The Profit is : ", Profit)
	// Using the Sprintf where u need'nt have to use the printf and println
	fmt.Print(RatioPrint)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculatevalues(Revenue, Expenses, TaxRate float64) (float64, float64, float64) {
	EBT := Revenue - Expenses
	Profit := EBT * (1 - TaxRate/100)
	Ratio := EBT / Profit

	return EBT, Profit, Ratio

}
