package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	conf := readConfig("config.json")

	files, err := os.ReadDir("jsonFiles")
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		data, err := os.ReadFile(fmt.Sprintf("jsonFiles/%s", file.Name()))
		if err != nil {
			panic(err)
		}
		decodedData := readJsonFile(data)
		for _, singleDataEntry := range decodedData {
			writeData(singleDataEntry, conf)
			time.Sleep(10 * time.Millisecond)
		}
		fmt.Printf("Dataset %d done\n", i+1)
	}

}
