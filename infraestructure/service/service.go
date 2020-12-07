package service

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/resty.v1"
)

// Service contain all the methods available to interact with external APIs
type Service interface {
	GetData() ([]byte, error)
}

type service struct {
	client *resty.Client
	url    string
}

// NewService will return a new instance of a service.
func NewService(url string) Service {
	client := resty.New()
	return service{client, url}
}

// GetData request all the available data through get method to the url inside Service structure.
func (s service) GetData() ([]byte, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
