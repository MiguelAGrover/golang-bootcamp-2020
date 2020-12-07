package controller

// AppController : contain all the controller interfaces allowed
type AppController struct {
	Digimon interface{ DigimonController }
}
