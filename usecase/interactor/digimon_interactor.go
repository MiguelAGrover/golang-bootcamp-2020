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

// DigimonInteractor contain the methods that the interactor should contain
type DigimonInteractor interface {
	Get(d []*model.Digimon) ([]*model.Digimon, error)
	Create(d *model.Digimon) (*model.Digimon, error)
	Update(d *model.Digimon) (*model.Digimon, error)
	Delete(d *model.Digimon) (*model.Digimon, error)
	GetSpecific(d *model.Digimon) (*model.Digimon, error)
}

// NewDigimonInteractor returns a digimon interactor based on the repository and a presenter.
func NewDigimonInteractor(r repository.DigimonRepository, p presenter.DigimonPresenter) DigimonInteractor {
	return &digimonInteractor{r, p}
}

// Get begin the process of retrieving all the digimon data as it is called from Input Port.
func (di *digimonInteractor) Get(d []*model.Digimon) ([]*model.Digimon, error) {
	d, err := di.DigimonRepository.FindAll(d)
	if err != nil {
		return nil, err
	}

	return di.DigimonPresenter.ResponseDigimons(d), nil
}

func (di *digimonInteractor) Create(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Create(d)

	return d, err
}

func (di *digimonInteractor) Update(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Update(d)

	return d, err
}

func (di *digimonInteractor) Delete(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Delete(d)

	return d, err
}

func (di *digimonInteractor) GetSpecific(d *model.Digimon) (*model.Digimon, error) {
	d, err := di.DigimonRepository.Find(d)

	return d, err
}
