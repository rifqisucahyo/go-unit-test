package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}
var categorySesrvice = CategorySesrvice{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	//Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categorySesrvice.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		ID:   "1",
		Name: "Laptop",
	}
	//Program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categorySesrvice.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}
