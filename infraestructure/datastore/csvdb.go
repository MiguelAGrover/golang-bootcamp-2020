package datastore

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

type CSVDB interface {
	LoadCSV() ([][]string, error)
	WriteCSV(rows []string) error
	WriteFullCSV(rows [][]string) error
	DropCSVFile() error
}

type csvDB struct {
	csvFile string
}

// LoadCSV return all the data that was retrieved from a csv
func (c csvDB) LoadCSV() ([][]string, error) {
	csvFile := fileExist(c.csvFile)
	defer csvFile.Close()
	records := readCSVFile(csvFile)
	if len(records) == 0 {
		return nil, &MyError{
			time.Now(),
			"Database is empty",
		}
	}
	return records, nil
}

func (c csvDB) WriteFullCSV(rows [][]string) error {
	if len(rows) == 0 {
		return &MyError{
			time.Now(),
			"Nothing to insert in arguments",
		}
	}
	csvFile := fileExist(c.csvFile)
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	err := writer.WriteAll(rows)
	if err != nil {
		return err
	}
	return nil
}

func (c csvDB) WriteCSV(row []string) error {
	csvFile := fileExist(c.csvFile)
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	writer.Write(row)
	writer.Flush()
	err := writer.Error()
	if err != nil {
		return err
	}
	return nil
}

func NewCSVDB(filepath string) csvDB {
	return csvDB{filepath}
}

func fileExist(filepath string) *os.File {
	var csvFile *os.File
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		csvFile = createFile(filepath)
	} else {
		csvFile = loadFile(filepath)
	}
	return csvFile
}

func loadFile(csvpath string) *os.File {
	csvFile, err := os.OpenFile(csvpath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	return csvFile
}

func createFile(csvpath string) *os.File {
	csvFile, err := os.Create(csvpath)
	if err != nil {
		log.Fatalln(err)
	}
	return csvFile
}

func readCSVFile(csvFile *os.File) [][]string {
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return records
}

func (c csvDB) DropCSVFile() error {
	err := os.Remove(c.csvFile)
	if !errors.Is(err, nil) {
		return err
	}
	return nil
}
