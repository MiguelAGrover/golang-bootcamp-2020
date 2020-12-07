package datastore

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// MyError : the structure that contain personalized error with information related to read and write to csvs
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s",
		e.What)
}

// CSVDB : the interface that contain the methods that are available for csv manipulation
type CSVDB interface {
	LoadCSV() ([][]string, error)
	WriteCSV(rows []string) error
	WriteFullCSV(rows [][]string) error
	DropCSVFile() error
	FileExist() bool
}

// csvDB : the structure that implements functions for CSVDB
type csvDB struct {
	csvFile string
}

// LoadCSV return all the data retrieved from the csv
func (c csvDB) LoadCSV() ([][]string, error) {
	csvFile := c.createOrOpenCSV(c.csvFile)
	defer csvFile.Close()
	records := c.readCSVFile(csvFile)
	if len(records) == 0 {
		return nil, &MyError{
			time.Now(),
			"Database is empty",
		}
	}
	return records, nil
}

// WriteFullCSV : Writes all the data passed as array to the csv
func (c csvDB) WriteFullCSV(rows [][]string) error {
	if len(rows) == 0 {
		return &MyError{
			time.Now(),
			"Nothing to write to CSV",
		}
	}
	csvFile := c.createOrOpenCSV(c.csvFile)
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	err := writer.WriteAll(rows)
	if err != nil {
		return err
	}
	return nil
}

// WriteCSV : Write one row to the csv
func (c csvDB) WriteCSV(row []string) error {
	if len(row) == 0 {
		return &MyError{
			time.Now(),
			"Nothing to write to CSV",
		}
	}
	csvFile := c.createOrOpenCSV(c.csvFile)
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

// FileExist : This verify that a file alreadtt exist
func (c csvDB) FileExist() bool {
	_, err := os.Stat(c.csvFile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// NewCSVDB : Return a structure of a new CSVDB given apath
func NewCSVDB(filepath string) CSVDB {
	return csvDB{filepath}
}

func (c csvDB) createOrOpenCSV(filepath string) *os.File {
	var csvFile *os.File
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		csvFile = c.createFile(filepath)
	} else {
		csvFile = c.loadFile(filepath)
	}
	return csvFile
}

func (c csvDB) loadFile(csvpath string) *os.File {
	csvFile, err := os.OpenFile(csvpath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	return csvFile
}

func (c csvDB) createFile(csvpath string) *os.File {
	csvFile, err := os.Create(csvpath)
	if err != nil {
		log.Fatalln(err)
	}
	return csvFile
}

func (c csvDB) readCSVFile(csvFile *os.File) [][]string {
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return records
}

// DropCSVFile : this delete the csv, this is should be used after data was stored in memory.
func (c csvDB) DropCSVFile() error {
	err := os.Remove(c.csvFile)
	if !errors.Is(err, nil) {
		return &MyError{
			time.Now(),
			"CSV not found",
		}
	}
	return nil
}
