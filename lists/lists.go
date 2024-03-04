package main

import "fmt"

func main() {
	precos := []float64{1, 2, 3}
	fmt.Println(precos[1])
	precos[1] = 9.99

	precos = append(precos, 5.99)
	precos = precos[1:]
	fmt.Println(precos)

	precosComDesconto := []float64{10, 50, 35}
	precos = append(precos, precosComDesconto...)
}

//func main() {
//	var produtos = [4]string{"Livro"}
//	precos := [4]float64{1.2, 3.4, 5.6, 5}
//	fmt.Println(precos)
//
//	produtos[2] = "Tapete"
//
//	fmt.Println(produtos)
//
//	fmt.Println(precos[2])
//
//	novosPrecos := precos[1:3]
//	//novosPrecos := precos[:3]
//	//novosPrecos := precos[1:]
//
//	fmt.Println(novosPrecos)
//}
