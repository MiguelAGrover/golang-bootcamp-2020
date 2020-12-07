package repository

import (
	"errors"
	"fmt"
	"wizegolangapi/domain/model"
	"wizegolangapi/infraestructure/datastore"
)

type digimonRepository struct {
	db datastore.CSVDB
}

type repositoryError struct {
	What string
}

func (e *repositoryError) Error() string {
	return fmt.Sprintf("%s",
		e.What)
}

// DigimonRepository : Shows all the methods to be implemented by a digimon repository
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
	Create(d *model.Digimon) (*model.Digimon, error)
	Update(d *model.Digimon) (*model.Digimon, error)
	Delete(d *model.Digimon) (*model.Digimon, error)
	Find(d *model.Digimon) (*model.Digimon, error)
}

// NewDigimonRepository : Returns an instance of a digimon repository.
func NewDigimonRepository(db datastore.CSVDB) DigimonRepository {
	return &digimonRepository{db}
}

// FindAll : Retrieve all the digimons from the database.
func (dr *digimonRepository) FindAll(d []*model.Digimon) ([]*model.Digimon, error) {
	data, err := dr.db.LoadCSV()
	if err != nil {
		return nil, err
	}
	var labels []string
	for index, rec := range data {
		if index == 0 {
			labels = rec
		} else {
			digimonmap := make(map[string]string)
			for i, value := range rec {
				digimonmap[labels[i]] = value
			}
			digimon := model.Digimon{
				Name:  digimonmap["Name"],
				Level: digimonmap["Level"],
				Image: digimonmap["Image"]}
			d = append(d, &digimon)
		}
	}

	return d, nil
}

// Create : Insert a digimon at the end of the csv.
func (dr *digimonRepository) Create(d *model.Digimon) (*model.Digimon, error) {

	if d == nil {
		return nil, &repositoryError{
			"No data was passed on Create",
		}
	}
	var DigimonString []string
	DigimonString = append(DigimonString, d.Name)
	DigimonString = append(DigimonString, d.Level)
	DigimonString = append(DigimonString, d.Image)

	if err := dr.db.WriteCSV(DigimonString); !errors.Is(err, nil) {
		return nil, err
	}

	return d, nil
}

// Update : Recreate the csv with the update requested.
func (dr *digimonRepository) Update(d *model.Digimon) (*model.Digimon, error) {
	data, err := dr.db.LoadCSV()

	if err != nil {
		return nil, err
	}

	var digimonArray []*model.Digimon

	for _, rec := range data {
		digimon := model.Digimon{Name: string(rec[0]), Level: string(rec[1]), Image: string(rec[2])}
		if d.Name == digimon.Name {
			if d.Level != "" {
				digimon.Level = d.Level
			}
			if d.Image != "" {
				digimon.Image = d.Image
			}
			d = &digimon
		}
		digimonArray = append(digimonArray, &digimon)
	}

	dr.db.DropCSVFile()
	var DigimonStringArray [][]string

	for _, digimon := range digimonArray {
		var row []string
		row = append(row, digimon.Name)
		row = append(row, digimon.Level)
		row = append(row, digimon.Image)
		DigimonStringArray = append(DigimonStringArray, row)
	}

	if err := dr.db.WriteFullCSV(DigimonStringArray); !errors.Is(err, nil) {
		return nil, err
	}

	return d, nil
}

// Delete : Delete the row of the csv based on the name of the digimon.
func (dr *digimonRepository) Delete(d *model.Digimon) (*model.Digimon, error) {
	data, err := dr.db.LoadCSV()

	if err != nil {
		return nil, err
	}

	var digimonArray []*model.Digimon

	for _, rec := range data {
		digimon := model.Digimon{Name: string(rec[0]), Level: string(rec[1]), Image: string(rec[2])}
		if d.Name != digimon.Name {
			digimonArray = append(digimonArray, &digimon)
		}
	}

	dr.db.DropCSVFile()
	var DigimonStringArray [][]string

	for _, digimon := range digimonArray {
		var row []string
		row = append(row, digimon.Name)
		row = append(row, digimon.Level)
		row = append(row, digimon.Image)
		DigimonStringArray = append(DigimonStringArray, row)
	}

	if err := dr.db.WriteFullCSV(DigimonStringArray); !errors.Is(err, nil) {
		return nil, err
	}

	return d, nil
}

// Find : Find a Dgigimon on the csv based on the name passed.
func (dr *digimonRepository) Find(d *model.Digimon) (*model.Digimon, error) {
	data, err := dr.db.LoadCSV()

	if !errors.Is(err, nil) {
		return nil, err
	}

	var digimon *model.Digimon

	var labels []string
	for index, rec := range data {
		if index == 0 {
			labels = rec
		} else {
			digimonmap := make(map[string]string)
			for i, value := range rec {
				digimonmap[labels[i]] = value
			}
			if digimonmap["Name"] == d.Name {
				digimon = new(model.Digimon)
				digimon.Name = digimonmap["Name"]
				digimon.Level = digimonmap["Level"]
				digimon.Image = digimonmap["Image"]
			}
		}
	}

	if digimon == nil {
		return nil, &repositoryError{
			"No digimon was found",
		}
	}

	return digimon, nil
}
