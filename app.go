package main

import (
	"fmt"

	"ybuilds.in/taxer/model"
	"ybuilds.in/taxer/util"
)

func main() {
	const inputPath string = "data/prices.txt"
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	done := make([]chan bool, len(taxRates))

	for i, rate := range taxRates {
		done[i] = make(chan bool)
		fileManager := util.NewFileManager(inputPath, fmt.Sprintf("%.2f", rate*100))
		job := model.NewTaxJob(rate, *fileManager)
		go job.Process(done[i])
	}

	for _, channel := range done {
		<-channel
	}
}
