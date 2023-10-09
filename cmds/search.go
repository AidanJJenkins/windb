package cmds

import (
	"encoding/json"
	"fmt"
	"os"
)

func GeneralSearch(tableName string, colName string, needle string) []string {
	file := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)
	fileData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	var data []map[string]string
	var res []string

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range data {
		if row[colName] == needle {
			var resultString string
			for key, value := range row {
				resultString += key + ": " + value + ", "
			}

			res = append(res, resultString)
		}
	}

	return res
}
