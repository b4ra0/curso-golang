package main

import "fmt"

type Produto struct {
	id     int
	titulo string
	preco  float64
}

func main() {
	//1)
	hobbies := [3]string{"Séries", "Futebol", "Videogame"}
	fmt.Println(hobbies)

	//2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3)
	listaDeHobbies := hobbies[:2]
	fmt.Println(listaDeHobbies)
	listaDeHobbies = []string{hobbies[0], hobbies[1]}
	fmt.Println(listaDeHobbies)

	//4)
	listaDeHobbies = hobbies[1:]
	fmt.Println(listaDeHobbies)

	//5)
	listaDeObjetivos := []string{"Ser efetivado no estágio", "Me formar"}
	fmt.Println(listaDeObjetivos)

	//6)
	listaDeObjetivos[1] = "Aprender a tocar bateria"
	listaDeObjetivos = append(listaDeObjetivos, "Organizar meu tempo")
	fmt.Println(listaDeObjetivos)

	//7)
	produto1 := Produto{
		id:     0,
		titulo: "Bateria Eletrônica",
		preco:  999.99,
	}
	produto2 := Produto{
		id:     1,
		titulo: "Guitarra",
		preco:  499.99,
	}
	listaDeProdutos := []Produto{produto1, produto2}
	fmt.Println(listaDeProdutos)

	produto3 := Produto{
		id:     2,
		titulo: "Microfone",
		preco:  199.99,
	}
	listaDeProdutos = append(listaDeProdutos, produto3)
	fmt.Println(listaDeProdutos)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
