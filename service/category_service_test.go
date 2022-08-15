package service

import (
	"testing"
	"unit-test-bayu/entity"
	"unit-test-bayu/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceMock_GetNotFound(t *testing.T) {
	// Mock program
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceMock_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Phone",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
