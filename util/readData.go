package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func (manager *FileManager) LoadData() ([]float64, error) {
	var prices []float64
	file, err := os.Open(manager.Input)

	if err != nil {
		fmt.Println("LoadData() - error opening file")
		fmt.Println(err)
		err = nil
		return nil, errors.New("LoadData() - error opening file")
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		value, err := strconv.ParseFloat(fileScanner.Text(), 64)

		if err != nil {
			fmt.Println("LoadData() - error parsing price")
			fmt.Println(err)
			err = nil
			return nil, errors.New("LoadData() - error parsing float")
		}

		prices = append(prices, value)
	}

	err = file.Close()

	if err != nil {
		fmt.Println("LoadData() - error closing file")
		fmt.Println(err)
		err = nil
		return nil, errors.New("LoadData() - error closing file")
	}

	return prices, nil
}
