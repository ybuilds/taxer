package model

import (
	"fmt"
	"strconv"

	"ybuilds.in/taxer/util"
)

type TaxJob struct {
	TaxRate     float64            `json:"tax_percent"`
	Prices      []float64          `json:"income"`
	TaxedPrices map[string]float64 `json:"taxed_income"`
	Manager     util.FileManager   `json:"-"`
}

func NewTaxJob(rate float64, fileManager util.FileManager) *TaxJob {
	return &TaxJob{
		TaxRate: rate,
		Prices:  []float64{10, 20, 30},
		Manager: fileManager,
	}
}

func (job *TaxJob) Process() {
	result := make(map[string]float64)
	prices, err := job.Manager.LoadData()

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
	err = job.Manager.WriteData(job)

	if err != nil {
		fmt.Println("Process() - error writing to file")
		fmt.Println(err)
		err = nil
		return
	}
}
