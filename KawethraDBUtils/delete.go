package kawethradbUtils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Delete(csvFilePath string, columnName string, columnValue interface{}) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	columnIndex := -1
	for i, name := range rows[0] {
		if name == columnName {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return fmt.Errorf("Column %s not found in CSV file", columnName)
	}

	var newRows [][]string
	for i, row := range rows {
		if i == 0 || row[columnIndex] != fmt.Sprintf("%v", columnValue) {
			newRows = append(newRows, row)
		}
	}

	file, err = os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range newRows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
