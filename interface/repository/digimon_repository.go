package repository

import (
	"errors"
	"wizegolangapi/domain/model"
	"wizegolangapi/infraestructure/datastore"
)

type digimonRepository struct {
	db datastore.CSVDB
}

// DigimonRepository shows all the methods to be implemented by a digimon repository
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
	Create(d *model.Digimon) (*model.Digimon, error)
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
