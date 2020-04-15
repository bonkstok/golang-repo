package main

import (
	"errors"
	"fmt"
)

const (
	goodCreditRate       = 450
	goodInterestRate     = 15.0
	goodMontlyPayRate    = 20.0
	defaultInterestRate  = 20.0
	defaultMontlyPayRate = 10.0
)

var (
	ErrorCreditScore = errors.New("Invalid credit score")
	ErrorIncome      = errors.New("Income invalid")
	ErrorLoanAmount  = errors.New("Loan amount invalid")
	ErrorLoanTerm    = errors.New("Loan term not a multiple of 12")
)

func loanChecker(creditScore int, income float32, loanAmount float32, loanTerm int) error {
	var (
		interestRate  float32
		maxMonthlyPay float32
		monthlyPay    float32
		interestTotal float32
		loanTotalCost float32
		approved      bool
	)
	if creditScore >= goodCreditRate {
		interestRate = goodInterestRate
		maxMonthlyPay = income * (goodMontlyPayRate / 100)
	} else if creditScore < 450 && creditScore > 0 {
		interestRate = defaultInterestRate
		maxMonthlyPay = income * (defaultMontlyPayRate / 100)
	} else {
		return ErrorCreditScore
	}

	if loanTerm%12 != 0 {
		return ErrorLoanTerm
	}

	interestTotal = loanAmount * (interestRate / 100)
	monthlyPay = interestTotal / float32(loanTerm)
	loanTotalCost = loanAmount + interestTotal
	//fmt.Println(interestRate, maxMontlyPay, interestTotal, monthlyPay)

	if creditScore > 0 && monthlyPay < maxMonthlyPay {
		approved = true
	}

	fmt.Println("Credit Score\t:", creditScore)
	fmt.Println("Income \t\t:", income)
	fmt.Println("Loan Amount \t:", loanAmount)
	fmt.Println("Loan Term \t:", loanTerm)
	fmt.Println("Max Monthly Pay :", maxMonthlyPay)
	fmt.Println("Monthly Pay \t:", monthlyPay)
	fmt.Println("Interest Rate \t:", interestRate)
	fmt.Println("Interest Total  :", interestTotal)
	fmt.Println("Total Cost \t:", loanTotalCost)
	fmt.Println("Approved \t:", approved)

	return ErrorCreditScore
}

func main() {
	loanChecker(300, 2500.0, 50000, 12) //creditscore, income, loadamount loanterm
}
