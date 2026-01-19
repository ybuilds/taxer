package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	taxes := make(map[float64][]float64)

	for _, rate := range taxRates {
		tax := make([]float64, len(prices))
		for index, price := range prices {
			tax[index] = price * (rate + 1)
		}
		taxes[rate] = tax
	}

	fmt.Println(taxes)
}
