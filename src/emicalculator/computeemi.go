package main

import "math"
import "fmt"

func CalculateEmi(emiRequest EmiRequest) []EquatedMonthlyInstallment {
	//PrintEmiRequest(emiRequest)
	time := emiRequest.MortgageTerm
	rate := emiRequest.AnnualCostOfMortgage / (12 * 100)
	amount := emiRequest.MortgageAmount
	
	emi := ComputeEmi(emiRequest)
	
	emiSlice := make([]EquatedMonthlyInstallment, time)

	for i := 0; i < time; i++ {

		interest := ComputeInterest(amount, rate)
		principal := emi - interest
		
		amount = amount - principal
		
		roundedPrincipal := roundOff(principal, 2)
		roundedInterest := roundOff(interest, 2)

		emiPayment := EquatedMonthlyInstallment{i, roundedPrincipal, roundedInterest, (roundedPrincipal+roundedInterest)}

		emiSlice[i] = emiPayment
	}

	return emiSlice
}

func PrintEmiRequest(emiRequest EmiRequest) {
		fmt.Println("Printing Received Request");
		fmt.Println(emiRequest.MortgageAmount)
		fmt.Println(emiRequest.AnnualCostOfMortgage)
		fmt.Println(emiRequest.MortgageTerm)
	}

func roundOff(number float64, places int) float64 {
		factor := math.Pow10(places)
			
		rounded := float64(int(number * factor))
		
		return rounded / factor 
	}

func ComputeInterest(amount, rate float64) float64 {
	return amount * rate
}

func ComputeEmi(emiRequest EmiRequest) float64 {

	rate := emiRequest.AnnualCostOfMortgage / (12 * 100)

	factor := math.Pow(1+rate, float64(emiRequest.MortgageTerm))

	numerator := emiRequest.MortgageAmount * rate * factor

	denominator := factor - 1

	emi := numerator / denominator
	
	return emi

}

