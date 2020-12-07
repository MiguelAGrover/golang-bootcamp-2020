package interactor

import (
	"wizegolangapi/domain/model"
	"wizegolangapi/usecase/presenter"
	"wizegolangapi/usecase/repository"
)

type digimonInteractor struct {
	DigimonRepository repository.DigimonRepository
	DigimonPresenter  presenter.DigimonPresenter
}

// DigimonInteractor : contain the methods that the interactor should contain.
type DigimonInteractor interface {
	Get(d []*model.Digimon) ([]*model.Digimon, error)
	Create(d *model.Digimon) (*model.Digimon, error)
	Update(d *model.Digimon) (*model.Digimon, error)
	Delete(d *model.Digimon) (*model.Digimon, error)
	GetSpecific(d *model.Digimon) (*model.Digimon, error)
}

// NewDigimonInteractor : returns a digimon interactor based on the repository and a presenter.
func NewDigimonInteractor(r repository.DigimonRepository, p presenter.DigimonPresenter) DigimonInteractor {
	return &digimonInteractor{r, p}
}

// Get : begin the process of retrieving all the digimon data as it is called from Input Port.
func (di *digimonInteractor) Get(d []*model.Digimon) ([]*model.Digimon, error) {
	d, err := di.DigimonRepository.FindAll(d)
	if err != nil {
		return nil, err
	}

	return di.DigimonPresenter.ResponseDigimons(d), nil
}

// Create : begin the process of creating a digimon into the csv as it is called from Input Port.
func (di *digimonInteractor) Create(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Create(d)

	return d, err
}

// Update : begin the process of updating a digimon into the csv as it is called from Input Port.
func (di *digimonInteractor) Update(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Update(d)

	return d, err
}

// Delete : begin the process of delete a digimon into the csv as it is called from Input Port.
func (di *digimonInteractor) Delete(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Delete(d)

	return d, err
}

// GetSpecified : begin the process of getting a digimon based on the name.
func (di *digimonInteractor) GetSpecific(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Find(d)

	return d, err
}
