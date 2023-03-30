package kawethradb

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func WriteCSV(filename string, records [][]string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	csvWriter.WriteAll(records)

	return nil
}

func Update(filename string, idColumnName string, id int, record []string) error {
	records, err := ReadCSV(filename)
	if err != nil {
		return err
	}

	var rowIndex int = -1
	for i, r := range records {
		if i == 0 {
			continue
		}
		if idValue, err := strconv.Atoi(r[indexOf(records[0], idColumnName)]); err == nil && idValue == id {
			rowIndex = i
			break
		}
	}

	if rowIndex != -1 {
		copy(records[rowIndex], record)
	} else {
		return fmt.Errorf("no record found with ID=%d", id)
	}

	if err := WriteCSV(filename, records); err != nil {
		return err
	}

	return nil
}

func indexOf(slice []string, target string) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1
}
