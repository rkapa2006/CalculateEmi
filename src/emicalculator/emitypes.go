package main

type EmiRequest struct {
	MortgageAmount       float64 //amount of mortgage or loan required
	AnnualCostOfMortgage float64 //annual % rate of interest
	MortgageTerm         int     //time in number of months needed to clear off the mortgage
}

type EmiRequestString struct {
	MortgageAmount       string //amount of mortgage or loan required
	AnnualCostOfMortgage string //annual % rate of interest
	MortgageTerm         string     //time in number of months needed to clear off the mortgage
}

type EquatedMonthlyInstallment struct {
	SerialNo       int
	PrincipalAmount float64
	InterestAmount   float64
	TotalAmount float64
}

