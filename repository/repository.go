package repository

import (
	"jaeger/model"
)

type ExampleRepository interface {
	Post(id *string) *model.Data
	Get() *model.Data
}

type exampleRepositoryDB struct {
}

func NewExampleRepository() ExampleRepository {
	return &exampleRepositoryDB{}
}

func (r *exampleRepositoryDB) Post(id *string) *model.Data {
	data := &model.Data{
		Data: *id,
	}

	return data
}

func (r *exampleRepositoryDB) Get() *model.Data {
	data := &model.Data{
		Data: "1",
	}

	return data
}
