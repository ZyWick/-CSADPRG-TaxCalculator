// ********************
// Last names: CABRERA, CORPUZ, GAMBOA, URETA
// Language: Go
// Paradigm(s): procedural programming
// ********************

package main

import (
	"fmt"
	"math"
)

// Formulas are ver 2022

func calcPhilHealth(monthlyIncome float64, ch chan float64) {
	ch <- math.Min(80000, math.Max(monthlyIncome, 10000)) * 0.04 / 2
	// halved since hati daw with employer
	// still 0.04 for 2023
}

func calcPagIbIG(monthlyIncome float64, ch chan float64) {
	rate := 0.01
	if monthlyIncome > 1500 {
		rate = 0.02
	}
	//diff for 2023
	ch <- math.Min(5000, monthlyIncome) * rate
}

func calcSSS(monthlyIncome float64, ch chan float64) {
	IncomeRange := 0
	switch {
	case monthlyIncome < 1000:
	case monthlyIncome < 3250:
		IncomeRange = 6
	case monthlyIncome > 24750:
		IncomeRange = 50
	default:
		IncomeRange = (int(monthlyIncome) + 250) / 500
	}

	ch <- float64(IncomeRange) * 500 * 0.0450
	//This is at 2023
}

func calcIncomeTax(TaxableIncome float64) float64 {
	var result float64
	switch {
	case TaxableIncome <= 20833:
		result = 0.00
	case TaxableIncome <= 33332:
		result = 0.00 + .2*(TaxableIncome-20833)
	case TaxableIncome <= 66666:
		result = 2500.00 + .25*(TaxableIncome-33333)
	case TaxableIncome <= 166666:
		result = 10833.33 + .3*(TaxableIncome-66667)
	case TaxableIncome <= 666666:
		result = 40833.33 + .32*(TaxableIncome-166667)
	default:
		result = 200833.33 + .35*(TaxableIncome-666667)
	}
	// 2023 is 15, 20, 25, 30, 35  for 2023
	return result
}

func main() {

	fmt.Println("\nPH Tax Calculator for 2022")
	fmt.Println("-----------------")
	fmt.Print("\n\nEnter Monthly Income: ")

	var monthlyIncome float64
	fmt.Scanln(&monthlyIncome)

	SSSch := make(chan float64)
	PhilHealthch := make(chan float64)
	PagIbIGch := make(chan float64)

	go calcSSS(monthlyIncome, SSSch)
	go calcPhilHealth(monthlyIncome, PhilHealthch)
	go calcPagIbIG(monthlyIncome, PagIbIGch)

	SSS := <-SSSch
	PhilHealth := <-PhilHealthch
	PagIBIG := <-PagIbIGch

	fmt.Printf("\nSSS: %.2f\n", SSS)
	fmt.Printf("PhilHealth: %.2f\n", PhilHealth)
	fmt.Printf("Pag-IBIG: %.2f\n", PagIBIG)

	TotalContributions := SSS + PhilHealth + PagIBIG
	TaxableIncome := monthlyIncome - TotalContributions

	IncomeTax := calcIncomeTax(TaxableIncome)
	TotalDeductions := IncomeTax + TotalContributions

	// fmt.Printf("\nTotal Contributions: %.2f\n", TotalContributions)
	// fmt.Printf("Taxable Income: %.2f\n", TaxableIncome)
	// fmt.Printf("Income Tax: %.2f\n", IncomeTax)

	fmt.Printf("\nTotal Deductions: %.2f\n", TotalDeductions)
	fmt.Printf("Net Monthly Pay: %.2f\n", monthlyIncome-TotalDeductions)

}
