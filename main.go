package kawethradb

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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

func UpdateByID(filename string, id int, record []string) error {
	records, err := ReadCSV(filename)
	if err != nil {
		return err
	}

	var rowIndex int = -1
	for i, r := range records {
		if i == 0 {
			continue
		}
		if idValue, err := strconv.Atoi(r[indexOf(records[0], "ID")]); err == nil && idValue == id {
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

		switch v := columnValue.(type) {
		case int:
			if val, err := strconv.Atoi(record[columnIndex]); err == nil && val == v {
				result := make(map[string]string)
				for i, name := range header {
					result[name] = record[i]
				}
				return result, nil
			}
		case float64:
			if val, err := strconv.ParseFloat(record[columnIndex], 64); err == nil && val == v {
				result := make(map[string]string)
				for i, name := range header {
					result[name] = record[i]
				}
				return result, nil
			}
		case string:
			if record[columnIndex] == v {
				result := make(map[string]string)
				for i, name := range header {
					result[name] = record[i]
				}
				return result, nil
			}
		default:
			return nil, fmt.Errorf("Unsupported column value type: %T", columnValue)
		}
	}

	return nil, fmt.Errorf("Csv dosyasında Kayıt bulunamadı %s = %v", columnName, columnValue)
}

func FindByID(csvFilePath string, columnValue interface{}) (map[string]string, error) {
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
		if name == "ID" {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return nil, fmt.Errorf("Column %s not found in CSV file", "ID")
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
	return nil, fmt.Errorf("Csv dosyasında Kayıt bulunamadı %s = %v", "ID", columnValue)
}

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

func DeleteByID(csvFilePath string, columnValue interface{}) error {
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
		if name == "ID" {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return fmt.Errorf("Column %s not found in CSV file", "ID")
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

func Count(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Dosya bulunamadı")
		return 0
	}
	defer file.Close()

	count, _ := csv.NewReader(file).ReadAll()
	newcount := len(count) - 1
	return newcount
}
