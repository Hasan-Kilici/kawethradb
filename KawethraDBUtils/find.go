package kawethradbUtils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Find(csvFilePath, columnName string, columnValue interface{}) (map[string]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	columnIndex := -1
	for i, name := range header {
		if name == columnName {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return nil, fmt.Errorf("Column %s not found in CSV file", columnName)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if record[columnIndex] == fmt.Sprintf("%v", columnValue) {
			result := make(map[string]string)
			for i, name := range header {
				result[name] = record[i]
			}
			return result, nil
		}
	}
	return nil, fmt.Errorf("Csv dosyasında Kayıt bulunamadı %s = %v", columnName, columnValue)
}
