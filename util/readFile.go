package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Read() ([]float64, error) {
	const path string = "data/income.txt"
	file, error := os.Open(path)

	incomes := []float64{}

	if error != nil {
		fmt.Println("Read() - Error reading data from file")
		return nil, error
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		income, error := strconv.ParseFloat(scanner.Text(), 64)

		if error != nil {
			fmt.Println("Read() - Error parsing float income from file")
			return nil, error
		}

		incomes = append(incomes, income)
	}

	return incomes, nil
}
