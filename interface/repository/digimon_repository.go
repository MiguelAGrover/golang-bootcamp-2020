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

// DigimonRepository shows all the methods to be implemented by a digimon repository
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
	Create(d *model.Digimon) (*model.Digimon, error)
	Update(d *model.Digimon) (*model.Digimon, error)
	Delete(d *model.Digimon) (*model.Digimon, error)
	Find(d *model.Digimon) (*model.Digimon, error)
}

// NewDigimonRepository Returns an instance of a digimon repository
func NewDigimonRepository(db datastore.CSVDB) DigimonRepository {
	return &digimonRepository{db}
}

// FindAll retrieve all the digimons from the database
func (dr *digimonRepository) FindAll(d []*model.Digimon) ([]*model.Digimon, error) {
	data, err := dr.db.LoadCSV()

	if err != nil {
		return nil, err
	}

	for _, rec := range data {
		digimon := model.Digimon{Name: string(rec[0]), Level: string(rec[1]), Image: string(rec[2])}
		d = append(d, &digimon)
	}

	return d, nil
}

func (dr *digimonRepository) Create(d *model.Digimon) (*model.Digimon, error) {

	var DigimonString []string
	DigimonString = append(DigimonString, d.Name)
	DigimonString = append(DigimonString, d.Level)
	DigimonString = append(DigimonString, d.Image)

	if err := dr.db.WriteCSV(DigimonString); !errors.Is(err, nil) {
		return nil, err
	}

	return d, nil
}

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

func (dr *digimonRepository) Find(d *model.Digimon) (*model.Digimon, error) {
	data, err := dr.db.LoadCSV()

	if !errors.Is(err, nil) {
		return nil, err
	}

	var digimon *model.Digimon

	for i, rec := range data {
		if i == 0 {
			fmt.Println(i)
		}
		if d.Name == string(rec[0]) {
			digimon = new(model.Digimon)
			digimon.Name = string(rec[0])
			digimon.Level = string(rec[1])
			digimon.Image = string(rec[2])
			break
		}
	}

	if digimon == nil {
		return nil, &repositoryError{
			"No digimon was found",
		}
	}

	return digimon, nil
}
