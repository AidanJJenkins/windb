package cmds

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrettyPrint(data []map[string]string) {
	for _, entry := range data {
		var row string
		for key, value := range entry {
			row += key + ": " + value + ", "
		}
		fmt.Println("Found: ")
		fmt.Println(row)
		fmt.Println("")
	}
}

func SearchJsonToArr(filePath string) ([]map[string]string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data []map[string]string

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
