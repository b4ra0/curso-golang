package main

import (
	"fmt"
	"math"
)

const taxaDeInflacao float64 = 6.5

func main() {
	var dinheiroInvestido, anos, taxaDeRetornoEsperado float64

	outputText("Informe o dinheiro investido: ", &dinheiroInvestido)
	outputText("Informe os anos: ", &anos)
	outputText("Informe a taxa de retorno esperado: ", &taxaDeRetornoEsperado)

	futureValue, futureRealValue := calcularFutureValue(dinheiroInvestido, taxaDeRetornoEsperado, anos)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (ajustado para inflação): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(mensagem string, teste *float64) {
	fmt.Print(mensagem)
	fmt.Scan(teste)
}

func calcularFutureValue(dinheiroInvestido, taxaDeRetornoEsperado, anos float64) (float64, float64) {
	fv := dinheiroInvestido * math.Pow(1+taxaDeRetornoEsperado/100, anos)
	rfv := fv / math.Pow(1+taxaDeInflacao/100, anos)
	return fv, rfv
}
