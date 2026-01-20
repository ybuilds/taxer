package main

import (
	"fmt"

	"ybuilds.in/taxer/model"
	"ybuilds.in/taxer/util"
)

func main() {
	const inputPath string = "data/prices.txt"
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, rate := range taxRates {
		fileManager := util.NewFileManager(inputPath, fmt.Sprintf("%.2f", rate*100))
		job := model.NewTaxJob(rate, *fileManager)
		job.Process()
	}
}
