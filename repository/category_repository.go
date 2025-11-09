package repository

import "basic-go/entity"

type CategoryRepository interface {
	FindByID(ID string) *entity.Category
}