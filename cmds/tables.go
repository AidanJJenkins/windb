package cmds

import (
	"encoding/json"
	"fmt"
	"os"
)

type Table struct {
	Name string   `json:"name"`
	Cols []string `json:"columns"`
}

func NewTable(name string) Table {
	return Table{name, nil}
}

func (T *Table) AddCol(n string) {
	T.Cols = append(T.Cols, n)
}

func WriteTable(table *Table, filename string) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	var data []*Table

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return err
	}

	data = append(data, table)

	updatedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	err = os.WriteFile(filename, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}

func CreateTableFile(name string) error {
	jsonArray := []map[string]interface{}{}

	jsonString, err := json.Marshal(jsonArray)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err

	}
	fileName := fmt.Sprintf("./tables/%s.json", name)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonString)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
