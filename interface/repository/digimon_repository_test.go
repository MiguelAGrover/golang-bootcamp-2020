package repository

import (
	"testing"
	"wizegolangapi/domain/model"
	"wizegolangapi/infraestructure/datastore"

	"github.com/stretchr/testify/assert"
)

func Test_CreateDigimon(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{"Test", "Test", "Test"}
	d, err := dr.Create(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_CreateEmptyDigimon(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	d, err := dr.Create(nil)
	assert := assert.New(t)
	assert.Nil(d)
	assert.NotNil(err)
}

func Test_DeleteDigimon(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{"Test", "Test", "Test"}
	d, err := dr.Delete(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_NotFoundDigimon(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{"NotFound", "", ""}
	d, err := dr.Find(digimon)
	assert := assert.New(t)
	assert.Nil(d)
	assert.NotNil(err)
}

func Test_FindDigimon(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{"Agumon", "", ""}
	d, err := dr.Find(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_UpdateDigimon(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	digimon := &model.Digimon{"Agumon", "Rookie", "https://digimon.shadowsmith.com/img/agumon.jpg"}
	d, err := dr.Update(digimon)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}

func Test_FindAllDigimons(t *testing.T) {
	db := datastore.NewCSVDB("../../test_utilities/test_data.csv")
	dr := NewDigimonRepository(db)
	d, err := dr.FindAll(nil)
	assert := assert.New(t)
	assert.NotNil(d)
	assert.Nil(err)
}
