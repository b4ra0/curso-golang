package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Revenue: ", &revenue)
	outputText("Expenses: ", &expenses)
	outputText("Tax Rate: ", &taxRate)

	ebt := calcularEBT(revenue, expenses)
	profit := calcularProfit(ebt, taxRate)
	ratio := calcularRatio(ebt, profit)

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}

func calcularRatio(ebt, profit float64) float64 {
	return ebt / profit
}

func calcularProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calcularEBT(revenue, expenses float64) float64 {
	return revenue - expenses
}

func outputText(mensagem string, endereco *float64) {
	fmt.Print(mensagem)
	fmt.Scan(endereco)
}
