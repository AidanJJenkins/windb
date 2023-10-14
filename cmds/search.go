package cmds

import (
	"fmt"
)

func GeneralSearch(tableName string, colName string, needle string) []map[string]string {
	file := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)
	var res []map[string]string

	data, err := SearchJsonToArr(file)
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range data {
		if row[colName] == needle {
			res = append(res, row)
		}
	}

	return res
}
