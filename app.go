package main

import (
	"fmt"

	"ybuilds.in/taxer/model"
	"ybuilds.in/taxer/util"
)

func main() {
	incomes, error := util.Read()

	if error != nil {
		fmt.Println("main() - Error fetching data from util")
	}

	taxRate := []float64{0, 7, 10, 15}

	for _, i := range taxRate {
		taxedSalary := model.NewSalary(incomes, i)
		taxedSalary.Process()
		taxedSalary.Save()
	}
}
