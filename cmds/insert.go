package cmds

import (
	"encoding/json"
	"fmt"
	"os"
)

// this should eventually check the index to make sure there are no repeats
// func GenId() int {
// id := rand.Intn(99999)
// check := tree.Find(id)
// if check != false {
// 	return GenId(tree)
// }
// return id
// }

func Insert(tableName string, columns []string, cData []string) error {
	fileName := fmt.Sprintf("./tables/%s.json", tableName)
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var data []map[string]string

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return err
	}

	entry := make(map[string]string)
	for i := 0; i < len(columns); i++ {
		entry[columns[i]] = cData[i]
	}
	// gI := GenId()
	// entry["id"] = strconv.Itoa(gI)

	data = append(data, entry)

	updatedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	err = os.WriteFile(fileName, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}

// get columns associated with the table
func InsertSetUp(name string, filename string) ([]string, error) {

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data []*Table
	var cols []string

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	for _, c := range data {
		if c.Name == name {
			cols = c.Cols
		}
	}

	return cols, nil
}
