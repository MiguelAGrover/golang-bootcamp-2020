package presenter

import (
	"wizegolangapi/domain/model"
	"wizegolangapi/usecase/presenter"
)

type digimonPresenter struct{}

// NewDigimonPresenter returns a pointer to DigimonPresenter
func NewDigimonPresenter() presenter.DigimonPresenter {
	return &digimonPresenter{}
}

// ResponseDigimons This method handle all the data before passing it to view.
func (dp *digimonPresenter) ResponseDigimons(ds []*model.Digimon) []*model.Digimon {
	return ds
}
