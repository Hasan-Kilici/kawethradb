package kawethradbUtils

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Insert(fileName string, records interface{}) error {
	value := reflect.ValueOf(records)
	if value.Kind() == reflect.Struct {
		return insertRecord(fileName, value)
	} else if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			if value.Index(i).Kind() != reflect.Struct {
				return fmt.Errorf("slice elemanları struct değil: %v", value.Index(i).Kind())
			}
			err := insertRecord(fileName, value.Index(i))
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return fmt.Errorf("records parametresi slice veya struct değil: %v", value.Kind())
	}
}

func insertRecord(fileName string, record reflect.Value) error {
	if record.Kind() != reflect.Struct {
		return fmt.Errorf("record parametresi struct değil: %v", record.Kind())
	}

	row := make([]string, 0, record.NumField())
	for i := 0; i < record.NumField(); i++ {
		field := record.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			row = append(row, strconv.FormatInt(field.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			row = append(row, strconv.FormatUint(field.Uint(), 10))
		case reflect.Float32, reflect.Float64:
			row = append(row, strconv.FormatFloat(field.Float(), 'f', -1, 64))
		case reflect.String:
			row = append(row, field.String())
		default:
			return fmt.Errorf("desteklenmeyen tip: %v", field.Kind())
		}
	}
	err := insertToFile(fileName, row, false)
	if err != nil {
		return err
	}

	return nil
}

func insertToFile(fileName string, row []string, isHeader bool) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if isHeader {
		headers := make([]string, 0, len(row))
		for _, val := range row {
			headers = append(headers, val)
		}
		if err := writer.Write(headers); err != nil {
			return err
		}
	}

	return writer.Write(row)
}
