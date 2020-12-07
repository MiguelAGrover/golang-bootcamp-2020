package presenter

import "wizegolangapi/domain/model"

// DigimonPresenter : this is the interface of the methods that digimon presenter has for showing data to the views
type DigimonPresenter interface {
	ResponseDigimons(d []*model.Digimon) []*model.Digimon
}
