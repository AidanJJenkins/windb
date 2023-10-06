package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aidanjjenkins/windb/cmds"
)

func TableCheck() {
	dir := "tables"
	file := "tables.json"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	filePath := dir + "/" + file

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
			err = cmds.WriteTable(&table, "./tables/tables.json")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Table Added")
		} else if strings.ToLower(input) == "insert" {
			var tName string

			fmt.Print("Enter name of table to insert into: ")
			scanner.Scan()
			tName = scanner.Text()
			fileName := fmt.Sprintf("./tables/%s.json", tName)

			var values []string

			columns, err := cmds.InsertSetUp(tName, "./tables/tables.json")
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
				inputFormat := fmt.Sprintf("Enter %s: ", c)

				var input string

				fmt.Print(inputFormat)
				scanner.Scan()
				input = scanner.Text()

				values = append(values, input)

			}

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
