package repository

import (
	"unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(Id string) *entity.Category {
	args := repository.Mock.Called(Id)
	if args.Get(0) == nil {
		return nil
	} else {

		category := args.Get(0).(entity.Category)
		return &category
	}

}
