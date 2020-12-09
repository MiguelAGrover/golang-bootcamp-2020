package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Service contain all the methods available to interact with external APIs
type Service interface {
	GetData() ([]byte, error)
}

// serviceError : error that will return if service fail
type serviceError struct {
	When time.Time
	What string
}

// Error will show messages using serviceError structure that comes from http request.
func (e *serviceError) Error() string {
	return fmt.Sprintf("%s",
		e.What)
}

type service struct {
	url string
}

// NewService will return a new instance of a service.
func NewService(url string) Service {
	return service{url}
}

// GetData request all the available data through get method to the url inside Service structure.
func (s service) GetData() ([]byte, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, &serviceError{time.Now(),
			"Request could not be done",
		}
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
