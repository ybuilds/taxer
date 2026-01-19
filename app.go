package main

import "ybuilds.in/taxer/model"

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, rate := range taxRates {
		job := model.NewTaxJob(rate)
		job.Process()
	}
}
