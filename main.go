package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("<My\\Path\\To\\Config> <My\\Path\\To\\JSON-Directory>")
		os.Exit(1)
	}

	configPath := os.Args[1]
	jsonPath := os.Args[2]

	conf := readConfig(configPath)

	files, err := os.ReadDir(jsonPath)
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", jsonPath, file.Name()))
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
