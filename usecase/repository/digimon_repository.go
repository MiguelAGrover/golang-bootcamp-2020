package repository

import "wizegolangapi/domain/model"

// DigimonRepository this is the interface that digimon repository should implement.
type DigimonRepository interface {
	FindAll(d []*model.Digimon) ([]*model.Digimon, error)
	Create(d *model.Digimon) (*model.Digimon, error)
	Update(d *model.Digimon) (*model.Digimon, error)
	Delete(d *model.Digimon) (*model.Digimon, error)
}
