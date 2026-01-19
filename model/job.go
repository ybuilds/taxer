package model

import (
	"fmt"
	"strconv"

	"ybuilds.in/taxer/util"
)

type TaxJob struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices map[string]float64
}

func NewTaxJob(rate float64) *TaxJob {
	return &TaxJob{
		TaxRate: rate,
		Prices:  []float64{10, 20, 30},
	}
}

func (job *TaxJob) Process() {
	result := make(map[string]float64)
	prices, err := util.LoadData()

	if err != nil {
		fmt.Println("Process() - error fetching prices")
		fmt.Println(err)
		err = nil
		return
	}

	job.Prices = prices

	for _, price := range job.Prices {
		value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", price*(job.TaxRate+1)), 64)

		if err != nil {
			fmt.Println("Process() - error parsing float")
			fmt.Println(err)
			err = nil
			return
		}

		result[fmt.Sprintf("%.2f", price)] = value
	}

	job.TaxedPrices = result
	err = util.WriteData(fmt.Sprintf("tax-%.0f", job.TaxRate*100), job)

	if err != nil {
		fmt.Println("Process() - error writing to file")
		fmt.Println(err)
		err = nil
		return
	}
}
