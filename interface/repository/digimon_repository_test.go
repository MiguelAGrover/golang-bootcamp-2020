package repository

import (
	"testing"

	"wizegolangapi/domain/model"
	"wizegolangapi/infraestructure/datastore"

	"github.com/stretchr/testify/assert"
)

const testingPath = "../../test_utilities/test_data.csv"
const testingField = "Test"

func Test_CreateDigimon(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{Name: testingField, Level: testingField, Image: testingField}
	d, err := dr.Create(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_CreateEmptyDigimon(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	d, err := dr.Create(nil)
	assert := assert.New(t)
	assert.Nil(d)
	assert.NotNil(err)
}

func Test_DeleteDigimon(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{Name: testingField, Level: testingField, Image: testingField}
	d, err := dr.Delete(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_NotFoundDigimon(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{Name: "NotFound", Level: "", Image: ""}
	d, err := dr.Find(digimon)
	assert := assert.New(t)
	assert.Nil(d)
	assert.NotNil(err)
}

func Test_FindDigimon(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{Name: "Agumon", Level: "", Image: ""}
	d, err := dr.Find(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_UpdateDigimon(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{Name: "Agumon", Level: "Rookie", Image: "https://digimon.shadowsmith.com/img/agumon.jpg"}
	d, err := dr.Update(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_FindAllDigimons(t *testing.T) {
	db := datastore.NewCSVDB(testingPath)
	dr := NewDigimonRepository(db)
	d, err := dr.FindAll(nil)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}
