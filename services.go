package main

import "strings"
import "errors"

type Service struct {
	Name string // name of the service
	File string // filepath for the plist
}

type Services struct {
	Services []Service
}

func (s *Services) GetNames() []string {
	var list []string
	for _, service := range s.Services {
		list = append(list, service.Name)
	}
	return list
}

func (s *Services) GetFiles() []string {
	var list []string
	for _, service := range s.Services {
		list = append(list, service.File)
	}
	return list
}

// Returns a filtered list of Services
func (s *Services) Filter(pattern string) Services {
	var list []Service
	var serv Services
	for _, service := range s.Services {
		if strings.Contains(service.Name, pattern) {
			list = append(list, service)
		}
	}
	serv.Services = list
	return serv
}

func (s *Services) Get(pattern string) (Service, error) {
	result := s.Filter(pattern)
	if len(result.Services) == 0 {
		return Service{}, nil
	} else if len(result.Services) > 1 {
		return Service{}, errors.New("More than one service returned")
	} else {
		return result.Services[0], nil
	}
}

// Adds an individual Service struct to the Services' list of Services
func (s *Services) AddService(service Service) {
	s.Services = append(s.Services, service)
}

// Returns the body of the plist file
func (s *Service) Show() string {
	return "body!"
}
