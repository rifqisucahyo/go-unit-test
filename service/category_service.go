package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type CategorySesrvice struct {
	Repository repository.CategoryRepository
}

func (service CategorySesrvice) Get(id string) (resp *entity.Category, err error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return resp, errors.New("Category not found")
	}

	return category, nil
}
