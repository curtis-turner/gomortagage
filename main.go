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

func calculatePayment(homeValue, downPayment, interestRate, pmi, homeInsurance, propertyTaxes float64, loanTerm int) float64 {
	fmt.Println(homeValue)
	fmt.Println(downPayment)
	var monthlyPayment float64
	var n = float64(loanTerm * 12)
	fmt.Println(n)
	var i = (interestRate / 100) / 12
	fmt.Println(i)
	var loanAmount = (homeValue - downPayment)
	fmt.Println(loanAmount)

	var discountFactor = ((math.Pow(1+i, n)) - 1) / (i * math.Pow(1+i, n))
	/*
		if pmiRequired(downPayment, homeValue) {
			fmt.Println("PMI Required")
			monthlyPayment = (loanAmount / discountFactor) + (pmi) + (homeInsurance) + (propertyTaxes)
		}
	*/

	monthlyPayment = (loanAmount / discountFactor) + (homeInsurance / 12) + (propertyTaxes / 12)

	return monthlyPayment
}

func main() {
	var homeValue float64
	fmt.Print("Enter the Home Value: ")
	fmt.Scan(&homeValue)
	var downPayment float64
	fmt.Print("Enter Your Down Payment: ")
	fmt.Scan(&downPayment)
	var interestRate float64
	fmt.Print("Enter Your Interest Rate: ")
	fmt.Scan(&interestRate)
	var loanTerm int
	fmt.Print("Enter Your Loan Term in Years: ")
	fmt.Scan(&loanTerm)
	var pmi float64
	fmt.Print("Enter Your PMI Percentage: ")
	fmt.Scan(&pmi)
	var homeInsurance float64
	fmt.Print("Enter Your Annual Home Insurance: ")
	fmt.Scan(&homeInsurance)
	var propertyTaxes float64
	fmt.Print("Enter Your Annual Property Taxes: ")
	fmt.Scan(&propertyTaxes)

	var monthlyPayment = calculatePayment(homeValue, downPayment, interestRate, pmi, homeInsurance, propertyTaxes, loanTerm)
	fmt.Printf("Your Monthly Payment is: $%f\n", monthlyPayment)
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
