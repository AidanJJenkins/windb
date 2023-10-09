package cmds

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func GetFromIndex(tableName string, id int) string {
	indexFile := fmt.Sprintf("./db/tables/%s/%sIndex.json", tableName, tableName)
	indexData, err := os.ReadFile(indexFile)
	if err != nil {
		return fmt.Sprintf("Error reading index file: %v", err)
	}

	var indexDataMap map[string]int
	err = json.Unmarshal(indexData, &indexDataMap)
	if err != nil {
		return fmt.Sprintf("Error parsing index data: %v", err)
	}

	index := indexDataMap[strconv.Itoa(id)]

	tableFile := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)
	tableData, err := os.ReadFile(tableFile)
	if err != nil {
		return fmt.Sprintf("Error reading table file: %v", err)
	}

	var tableDataArray []map[string]string
	err = json.Unmarshal(tableData, &tableDataArray)
	if err != nil {
		return fmt.Sprintf("Error parsing table data: %v", err)
	}

	// index, err := strconv.Atoi(location)
	if err != nil {
		return "Invalid location in index"
	}

	result := tableDataArray[index]
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("Error encoding result to JSON: %v", err)
	}

	return string(resultJSON)
}
