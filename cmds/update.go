package cmds

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func Update(tableName, colName string, changed string, id int) error {
	toBeUpdated, idx, err := findId(id, tableName)
	if err != nil || idx == -1 {
		return err
	}

	file := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)
	data, err := SearchJsonToArr(file)
	if err != nil {
		return err
	}

	toBeUpdated[colName] = changed

	data[idx] = toBeUpdated

	updatedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	err = os.WriteFile(file, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

// should also return index to use in udpate
func findId(id int, tableName string) (map[string]string, int, error) {
	indexFile := fmt.Sprintf("./db/tables/%s/%sIndex.json", tableName, tableName)
	indexData, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, -1, err
	}

	var indexDataArray []map[string]int
	err = json.Unmarshal(indexData, &indexDataArray)
	if err != nil {
		return nil, -1, err
	}

	if len(indexDataArray) == 0 {
		return nil, -1, err
	}

	index := indexDataArray[0][strconv.Itoa(id)]
	fmt.Println("index: ", index)

	tableFile := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)
	tableData, err := os.ReadFile(tableFile)
	if err != nil {
		return nil, -1, err
	}

	var tableDataArray []map[string]string
	err = json.Unmarshal(tableData, &tableDataArray)
	if err != nil {
		return nil, -1, err
	}

	if len(indexDataArray) == 0 {
		return nil, -1, err
	}

	if index < 0 || index >= len(tableDataArray) {
		return nil, -1, err
	}

	result := tableDataArray[index]

	return result, index, nil
}
