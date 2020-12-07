package controller

import (
	"errors"
	"net/http"

	"wizegolangapi/domain/model"
	"wizegolangapi/usecase/interactor"
)

type digimonController struct {
	digimonInteractor interactor.DigimonInteractor
}

// DigimonController : This interface will handle the request for Digimons that comes from outer layers
type DigimonController interface {
	GetDigimons(c Context) error
	CreateDigimon(c Context) error
	UpdateDigimon(c Context) error
	DeleteDigimon(c Context) error
	GetDigimon(c Context) error
}

// NewDigimonController : This function returns a Digimon controller based on the interactor which catch the request for digimon data
func NewDigimonController(di interactor.DigimonInteractor) DigimonController {
	return &digimonController{di}
}

// GetDigimons retrieves data from Digimons and return response as json or shows and error.
func (dc *digimonController) GetDigimons(c Context) error {
	var d []*model.Digimon

	d, err := dc.digimonInteractor.Get(d)
	if err != nil {
		return c.JSON(http.StatusNotFound, d)
	}

	return c.JSON(http.StatusOK, d)
}

// CreateDigimon receives the data through a form and insert data to the csv.
func (dc *digimonController) CreateDigimon(c Context) error {
	var params model.Digimon

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	d, err := dc.digimonInteractor.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusCreated, d)
}

// UpdateDigimon : This will call the update method of a digimon based on the received form
func (dc *digimonController) UpdateDigimon(c Context) error {
	var params model.Digimon

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	d, err := dc.digimonInteractor.Update(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusAccepted, d)
}

// DeleteDigimon : This will call the delete digimon based on the name received in the URL
func (dc *digimonController) DeleteDigimon(c Context) error {
	name := c.Param("name")

	params := model.Digimon{Name: name}

	d, err := dc.digimonInteractor.Delete(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusAccepted, d)
}

// GetDigimon : This will call the get digimon based on the name received in the URL
func (dc *digimonController) GetDigimon(c Context) error {
	name := c.Param("name")
	params := model.Digimon{Name: name}

	d, err := dc.digimonInteractor.GetSpecific(&params)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, d)
}
