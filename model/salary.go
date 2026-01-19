package model

import (
	"encoding/json"
	"fmt"
	"os"

	"ybuilds.in/taxer/util"
)

type Salary struct {
	InputSalary []float64
	TaxRate     float64
	TaxedSalary map[string]float64
}

func NewSalary(input []float64, taxRate float64) *Salary {
	inputSalary, error := util.Read()

	if error != nil {
		fmt.Println("NewSalary() - Error fetching data from util")
		return nil
	}

	return &Salary{
		inputSalary,
		taxRate,
		nil,
	}
}

func (salary *Salary) Process() {
	res := make(map[string]float64)

	for _, i := range salary.InputSalary {
		res[fmt.Sprintf("%.2f", i)] = i * (1 + (salary.TaxRate / 100))
	}

	fmt.Println(res)
}

func (salary *Salary) Save() error {
	const path string = "data/tax.json"

	data, error := json.Marshal(salary.TaxedSalary)
	fmt.Printf("Data: %v", data)

	if error != nil {
		fmt.Println("save() - Error marshalling salary map")
		return error
	}

	os.WriteFile(path, data, 0644)

	return nil
}
