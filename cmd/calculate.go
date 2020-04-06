/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// Property represents property data generated from an MLS listing URL
type Property struct {
	URL            string  `json:"url"` // flag
	HomeValue      float64 // pulled from scraper
	DownPayment    float64 `json:"downPayment"` // flag
	LoanAmount     float64 // generated in calculate function
	InterestRate   float64 `json:"interestRate"`  // flag
	LoanTerm       int64   `json:"loanTerm"`      // flag
	PMI            float64 `json:"pmi"`           // flag
	HomeInsurance  float64 `json:"homeInsurance"` // flag
	PropertyTaxes  float64 // pulled from scraper
	HOA            float64 // pullfed from scraper
	MonthlyPayment float64 // generated by the calculate function
}

func Scrape(p *Property) {
	doc, err := goquery.NewDocument(p.URL)
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("span").Each(func(index int, span *goquery.Selection) {
		spanText := span.Text()
		if spanText == "List Price:" {
			if s := span.Next(); s != nil {
				fmt.Println("Home Value: ", s.Text())
				var value = strings.Replace(s.Text(), "$", "", -1)
				value = strings.Replace(value, ",", "", -1)
				if p.HomeValue, err = strconv.ParseFloat(value, 64); err != nil {
					fmt.Println(err)
				}
			}
		}
	})

	// get the Annual HOA Fees and Property Taxes
	doc.Find("div").Each(func(index int, div *goquery.Selection) {
		div.Find("span").Each(func(index int, span *goquery.Selection) {
			if span.Text() == "Tax Annual Amount" {
				if d := div.Next(); d != nil {
					d.Find("span").Each(func(index int, span *goquery.Selection) {
						if strings.Contains(span.Text(), "$") {
							fmt.Println("Annual Property Taxes: ", span.Text())
							var value = strings.Replace(span.Text(), "$", "", -1)
							value = strings.Replace(value, ",", "", -1)
							if p.PropertyTaxes, err = strconv.ParseFloat(value, 64); err != nil {
								fmt.Println(err)
							}
						}
					})
				}
			}
			if span.Text() == "Total Annual HOA Fees" {
				if d := div.Next(); d != nil {
					d.Find("span").Each(func(index int, span *goquery.Selection) {
						if strings.Contains(span.Text(), "$") {
							fmt.Println("Annual HOA Amount: ", span.Text())
							var value = strings.Replace(span.Text(), "$", "", -1)
							value = strings.Replace(value, ",", "", -1)
							if p.HOA, err = strconv.ParseFloat(value, 64); err != nil {
								fmt.Println(err)
							}
						}
					})
				}
			}
		})
	})
}

func MonthlyPayment(p *Property) {

	var n = float64(p.LoanTerm * 12)
	var i = (p.InterestRate / 100) / 12
	var loanAmount = (p.HomeValue - p.DownPayment)
	p.LoanAmount = loanAmount

	var discountFactor = ((math.Pow(1+i, n)) - 1) / (i * math.Pow(1+i, n))

	p.MonthlyPayment = (loanAmount / discountFactor) + (p.HomeInsurance / 12) + (p.PropertyTaxes / 12) + (p.HOA / 12)

	if pmiRequired(p) {
		fmt.Println("PMI Required")
		var pmi = getPmi(p)
		p.MonthlyPayment = (loanAmount / discountFactor) + (pmi / 12) + (p.HomeInsurance / 12) + (p.PropertyTaxes / 12) + (p.HOA / 12)
	}

	fmt.Println("Your Monthly Payment is: $", fmt.Sprintf("%.2f", p.MonthlyPayment))
}

func pmiRequired(p *Property) bool {
	if ((p.DownPayment / p.HomeValue) * 100) < 20.00 {
		return true
	}
	return false
}

func getPmi(p *Property) float64 {
	var annualPmi float64
	annualPmi = p.LoanAmount * (p.PMI / 100)
	return annualPmi
}

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calcuate Monthly Mortgage Payment based on Zillow URL given.",
	Long: `Calcuate Monthly Mortgage Payment based on Zillow URL given.
	For Example:
	gomortgage calcualte https://www.zillow.com/homedetails/4294-N-Meadows-Dr-Castle-Rock-CO-80109/124172826_zpid/`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calculating Monthly Payment")
		var property Property
		property.URL, _ = cmd.Flags().GetString("url")
		property.DownPayment, _ = cmd.Flags().GetFloat64("down-payment")
		property.InterestRate, _ = cmd.Flags().GetFloat64("interest-rate")
		property.LoanTerm, _ = cmd.Flags().GetInt64("loan-term")
		property.PMI, _ = cmd.Flags().GetFloat64("pmi")
		property.HomeInsurance, _ = cmd.Flags().GetFloat64("home-insurance")
		Scrape(&property)
		MonthlyPayment(&property)
	},
}

// variables for flags
var (
	url           string  // flag
	downPayment   float64 // flag
	interestRate  float64 // flag
	loanTerm      int64   // flag
	pmi           float64
	homeInsurance float64 // flag
)

func init() {

	rootCmd.AddCommand(calculateCmd)
	calculateCmd.Flags().StringVar(&url, "url", "", "MLS Listing URL to Scrape data from")
	calculateCmd.Flags().Float64Var(&downPayment, "down-payment", 14000.00, "Down Payment Amount")
	calculateCmd.Flags().Float64Var(&interestRate, "interest-rate", 3.2, "Annual Interest Rate")
	calculateCmd.Flags().Int64Var(&loanTerm, "loan-term", 30, "Number of Years for the Loan Term")
	calculateCmd.Flags().Float64Var(&pmi, "pmi", 0.85, "PMI Rate for the Loan")
	calculateCmd.Flags().Float64Var(&homeInsurance, "home-insurance", 1700.00, "Annual Home Insurance Premium")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
