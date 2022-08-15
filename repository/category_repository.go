package repository

import "unit-test-bayu/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
