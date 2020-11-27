package registry

import (
	"wizegolangapi/interface/controller"
	ip "wizegolangapi/interface/presenter"
	ir "wizegolangapi/interface/repository"
	"wizegolangapi/usecase/interactor"
	up "wizegolangapi/usecase/presenter"
	ur "wizegolangapi/usecase/repository"
)

// NewDigimonController Returns a controller with interactor for digimons
func (r *registry) NewDigimonController() controller.DigimonController {
	return controller.NewDigimonController(r.NewDigimonInteractor())
}

// NewDigimonInteractor Returns an interactor with a repository and a presenter
func (r *registry) NewDigimonInteractor() interactor.DigimonInteractor {
	return interactor.NewDigimonInteractor(r.NewDigimonRepository(), r.NewDigimonPresenter())
}

// NewDigimonRepository returns a repository of interface that will be used to pass to lower layers the database
func (r *registry) NewDigimonRepository() ur.DigimonRepository {
	return ir.NewDigimonRepository(r.db)
}

// NewDigimonPresenter returns a presenter for digimons
func (r *registry) NewDigimonPresenter() up.DigimonPresenter {
	return ip.NewDigimonPresenter()
}
