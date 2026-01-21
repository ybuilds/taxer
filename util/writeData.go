package util

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func (manager *FileManager) WriteData(data any) error {
	var path string = fmt.Sprintf("data/output/tax_%v.json", manager.Output)
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("WriteData() - error creating/opening file")
		fmt.Println(err)
		err = nil
		return err
	}

	time.Sleep(time.Second * 5)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		fmt.Println("WriteData() - error encoding data to json")
		fmt.Println(err)
		err = nil
		return err
	}

	err = file.Close()

	if err != nil {
		fmt.Println("WriteData() - error closing file")
		fmt.Println(err)
		err = nil
		return err
	}

	return nil
}
