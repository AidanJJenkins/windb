package cmds

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func SearchAndDestroy(tableName string, colName string, needle string) error {
	if colName == "id" {
		fmt.Println("delete by id")
	} else {
		file := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)

		data, err := SearchJsonToArr(file)
		if err != nil {
			return err
		}

		for i, row := range data {
			if row[colName] == needle {
				id, err := strconv.Atoi(row["id"])
				if err != nil {
					return err
				}
				// slice entry from range, appending each side back together
				data = append(data[:i], data[i+1:]...)
				err = idDelete(id, tableName)
				if err != nil {
					return err
				}

				fmt.Println(i)
				break
			}
		}
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
	}
	return nil
}

func idDelete(id int, tableName string) error {
	indexFile := fmt.Sprintf("./db/tables/%s/%sIndex.json", tableName, tableName)
	indexData, err := os.ReadFile(indexFile)
	if err != nil {
		return err
	}

	var indexDataArray []map[string]int
	err = json.Unmarshal(indexData, &indexDataArray)
	if err != nil {
		return err
	}

	if len(indexDataArray) == 0 {
		fmt.Println("Empty data")
	}

	i := indexDataArray[0][strconv.Itoa(id)]

	file := fmt.Sprintf("./db/tables/%s/%s.json", tableName, tableName)
	data, err := SearchJsonToArr(file)
	if err != nil {
		return err
	}

	if data[i]["id"] == strconv.Itoa(id) {
		data = append(data[:i], data[i+1:]...)
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
	}

	delete(indexDataArray[0], strconv.Itoa(id))

	updatedJSONMap, err := json.Marshal(indexData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	err = os.WriteFile(file, updatedJSONMap, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}
