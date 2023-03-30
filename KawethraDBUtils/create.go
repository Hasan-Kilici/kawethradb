package kawethradbutils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func CreateDB(dbName, fileName string, records interface{}) error {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("Veritabanı %s okunuyor\n", fileName)
		return nil
	}

	value := reflect.ValueOf(records)
	if value.Kind() != reflect.Slice {
		return fmt.Errorf("records parametresi bir slice değil: %v", value.Kind())
	}
	slice := value.Interface()

	fmt.Println(slice)

	if value.Len() == 0 {
		return errors.New("records parametresi boş bir slice")
	}
	first := value.Index(0)
	if first.Kind() != reflect.Struct {
		return fmt.Errorf("records parametresinin elemanları struct değil: %v", first.Kind())
	}

	headers := make([]string, 0, first.NumField())
	for i := 0; i < first.NumField(); i++ {
		headers = append(headers, first.Type().Field(i).Name)
	}
	err := writeToFile(fileName, headers)
	if err != nil {
		return err
	}

	for i := 0; i < value.Len(); i++ {
		row := make([]string, 0, first.NumField())
		for j := 0; j < first.NumField(); j++ {
			field := value.Index(i).Field(j)
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
				return fmt.Errorf("Desteklenmeyen veri tipi: %v", field.Kind())
			}
		}
		err = writeToFile(fileName, row)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeToFile(fileName string, row []string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write(row)
}
