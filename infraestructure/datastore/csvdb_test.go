package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testingPath = "../../test_utilities/test_data.csv"
const testingPath2 = "../../test_utilities/test_data2.csv"
const testingPath3 = "../../test_utilities/test_data3.csv"

func Test_ReadCSV(t *testing.T) {
	db := NewCSVDB(testingPath)
	data, _ := db.LoadCSV()
	assert := assert.New(t)
	assert.GreaterOrEqual(len(data), 2)
}

func Test_ReadEmptyCSV(t *testing.T) {
	db := NewCSVDB(testingPath2)
	data, err := db.LoadCSV()
	assert := assert.New(t)
	assert.EqualError(err, "Database is empty")
	assert.Equal(0, len(data))
}

func Test_WriteCSV(t *testing.T) {
	db := NewCSVDB(testingPath)
	newdata := []string{"Name", "Level", "Image"}
	err := db.WriteCSV(newdata)
	assert.Nil(t, err)
}

func Test_WriteNothingToCSV(t *testing.T) {
	db := NewCSVDB(testingPath)
	newdata := []string{}
	err := db.WriteCSV(newdata)
	assert := assert.New(t)
	assert.EqualError(err, "Nothing to write to CSV")
}

func Test_WriteEmptyRowsCSV(t *testing.T) {
	db := NewCSVDB(testingPath)
	newdata := [][]string{}
	err := db.WriteFullCSV(newdata)
	assert := assert.New(t)
	assert.EqualError(err, "Nothing to write to CSV")
}

func Test_WriteRowsCSV(t *testing.T) {
	db := NewCSVDB(testingPath)
	newdata := [][]string{{"Example", "Example", "Example"}}
	err := db.WriteFullCSV(newdata)
	assert.Nil(t, err)
}

func Test_DropCSVFile(t *testing.T) {
	db := NewCSVDB(testingPath2)
	err := db.DropCSVFile()
	assert.Nil(t, err)
}

func Test_DropNonExistentCSVFile(t *testing.T) {
	db := NewCSVDB(testingPath3)
	err := db.DropCSVFile()
	assert.EqualError(t, err, "CSV not found")
}
