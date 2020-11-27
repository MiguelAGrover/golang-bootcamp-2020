package registry

import (
	"wizegolangapi/infraestructure/datastore"
	"wizegolangapi/interface/controller"
)

type registry struct {
	db datastore.CSVDB
}

// Registry contain a pointer to an instance of form DB
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry returns a new registry with a pointer to the database.
func NewRegistry(db datastore.CSVDB) Registry {
	return &registry{db}
}

// NewAppController Creates an app controller that includes the controller for other models.
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Digimon: r.NewDigimonController(),
	}
}
