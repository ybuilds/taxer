package main

import (
	"fmt"

	"ybuilds.in/taxer/util"
)

func main() {
	incomes, error := util.Read()
	if error != nil {
		fmt.Println("main() - Error reading data from file")
		return
	}

	fmt.Println(incomes)
}
