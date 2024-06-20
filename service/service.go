package service

import (
	"jaeger/model"
	"jaeger/repository"
)

type ExampleService interface {
	Post(id *string) *model.Data
	Get() *model.Data
}

type exampleService struct {
	exampleRepo repository.ExampleRepository
}

func NewExampleService(exampleRepo repository.ExampleRepository) ExampleService {
	return &exampleService{exampleRepo: exampleRepo}
}

func (s *exampleService) Post(id *string) *model.Data {
	response := s.exampleRepo.Post(id)

	return response
}

func (s *exampleService) Get() *model.Data {
	response := s.exampleRepo.Get()

	return response
}
