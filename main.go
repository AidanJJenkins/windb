package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aidanjjenkins/windb/cmds"
)

func TableCheck() {
	dir := "db"
	info := "info"
	file := "tableInfo.json"
	tablesDir := "tables"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(dir + "/" + info); os.IsNotExist(err) {
		if err := os.Mkdir(dir+"/"+info, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(dir + "/" + tablesDir); os.IsNotExist(err) {
		if err := os.Mkdir(dir+"/"+tablesDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	filePath := dir + "/" + info + "/" + file

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fileHandle, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer fileHandle.Close()
		jsonArray := []map[string]interface{}{}
		jsonString, err := json.Marshal(jsonArray)

		_, err = fileHandle.Write([]byte(jsonString))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		TableCheck()

		fmt.Print("Enter a command or type 'quit' to quit: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "quit" {
			fmt.Println("Quitting session.")
			break
		} else if strings.ToLower(input) == "add table" {
			var tN string
			adding := true

			fmt.Print("Enter table name: ")
			scanner.Scan()
			tN = scanner.Text()

			err := cmds.CreateTableFile(tN)
			if err != nil {
				fmt.Println(err)
			}

			table := cmds.NewTable(tN)
			table.AddCol("id")

			for adding {
				var cN string

				fmt.Print("Enter column name: ")
				scanner.Scan()
				cN = scanner.Text()
				if cN == "q" {
					break
				}
				table.AddCol(cN)
			}
			err = cmds.WriteTable(&table, "./db/info/tableInfo.json")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Table Added")
		} else if strings.ToLower(input) == "insert" {
			var tName string

			fmt.Print("Enter name of table to insert into: ")
			scanner.Scan()
			tName = scanner.Text()
			fileName := fmt.Sprintf("./db/tables/%s.json", tName)

			var values []string

			columns, err := cmds.InsertSetUp(tName, "./db/info/tableInfo.json")
			if err != nil {
				fmt.Println(err)
			}

			fileData, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
			}

			var data [][]map[string]string
			err = json.Unmarshal(fileData, &data)
			if err != nil {
				fmt.Println(err)
			}

			for _, c := range columns {
				if c == "id" {
					idx := cmds.GenId()
					values = append(values, strconv.Itoa(idx))
					continue
				}
				inputFormat := fmt.Sprintf("Enter %s: ", c)

				var input string

				fmt.Print(inputFormat)
				scanner.Scan()
				input = scanner.Text()

				values = append(values, input)
			}

			// fmt.Println(values)
			cmds.Insert(tName, columns, values)
		} else {
			badInput := strings.ToLower(input)
			fmt.Printf("'%s' is not a valid command, please retry\n", badInput)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func uQuit(input string) bool {
	if input == "quit" {
		return true
	}
	return false
}
