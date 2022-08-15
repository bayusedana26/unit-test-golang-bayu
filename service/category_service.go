package service

import (
	"errors"
	"unit-test-bayu/entity"
	"unit-test-bayu/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
