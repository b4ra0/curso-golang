package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers, double)

	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int, trasform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, double(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
