package main

import (
	"fmt"
	"math"
)

func pmiRequired(downPayment, homeValue float64) bool {
	if ((downPayment / homeValue) * 100) < 20.00 {
		return true
	}
	return false
}

func calculatePmi(homeValue, loanAmount, pmi float64) float64 {
	var monthlyPmi float64
	return monthlyPmi
}

func calculatePayment(homeValue, downPayment, loanTerm, interestRate, pmi, homeInsurance, propertyTaxes float64) float64 {
	var monthlyPayment float64
	var n = loanTerm * 12
	var i = interestRate / 100
	var loanAmount = homeValue - downPayment

	var discountFactor = ((math.Pow(1+i, n)) - 1) / (i * math.Pow(1+i, n))
	if pmiRequired(downPayment, homeValue) {
		monthlyPayment = (loanAmount / discountFactor) + (pmi) + (homeInsurance) + (propertyTaxes)
	}

	monthlyPayment = (loanAmount / discountFactor) + (homeInsurance / 12) + (propertyTaxes / 12)

	return monthlyPayment
}

func main() {
	var homeValue = 0.00
	fmt.Print("Enter the Home Value: ")
	fmt.Scan(&homeValue)
	fmt.Println(homeValue)
	var downPayment = 0.00
	fmt.Print("Enter Your Down Payment: ")
	fmt.Scan(&downPayment)
	fmt.Println(downPayment)
	var interestRate = 0.00
	fmt.Print("Enter Your Interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Println(interestRate)
	var loanTerm = 0
	fmt.Print("Enter Your Loan Term in Years: ")
	fmt.Scan(&loanTerm)
	fmt.Println(loanTerm)
	var pmi = 0.00
	fmt.Print("Enter Your PMI Percentage: ")
	fmt.Scan(&pmi)
	fmt.Println(pmi)
	var homeInsurance = 0
	fmt.Print("Enter Your Annual Home Insurance: ")
	fmt.Scan(&homeInsurance)
	fmt.Println(homeInsurance)
	var propertyTaxes = 0
	fmt.Print("Enter Your Annual Property Taxes: ")
	fmt.Scan(&propertyTaxes)
	fmt.Println(propertyTaxes)

	var monthlyPayment = calculatePayment(homeInsurance, downPayment, loanTerm, interestRate, pmi, homeInsurance, propertyTaxes)

}

/*
var homeValue float64
	var downPayment float64
	var interestRate float64
	var loanTerm float64
	var pmi float64
	var homeInsurance float64
	var propertyTaxes float64

	//read inputs from user
	fmt.Print("Enter Home Value: ")
	if _, err := fmt.Scanf("%f", &homeValue); err == nil {
		fmt.Printf("Home Value Entered: %f\n", homeValue)
	}

	fmt.Print("Enter Down Payment: ")
	if _, err := fmt.Scanf("%f", &downPayment); err == nil {
		fmt.Printf("Down Payment Entered: %f\n", downPayment)
	}

	fmt.Print("Enter Interest Rate: ")
	if _, err := fmt.Scanf("%f", &interestRate); err == nil {
		fmt.Printf("Interest Rate Entered: %f\n", interestRate)
	}

	fmt.Print("Enter Loan Term in Years: ")
	if _, err := fmt.Scanf("%f", &loanTerm); err == nil {
		fmt.Printf("Loan Term Entered: %f\n", loanTerm)
	}

	fmt.Print("Enter PMI as percentage: ")
	if _, err := fmt.Scanf("%f", &pmi); err == nil {
		fmt.Printf("PMI Entered: %f\n", pmi)
	}

	fmt.Print("Enter Annual Home Insurance: ")
	if _, err := fmt.Scanf("%f", &homeInsurance); err == nil {
		fmt.Printf("Home Insurance Entered: %f\n", homeInsurance)
	}

	fmt.Print("Enter Annual Property Taxes: ")
	if _, err := fmt.Scanf("%f", &propertyTaxes); err == nil {
		fmt.Printf("Property Taxes Entered: %f\n", propertyTaxes)
	}
*/
