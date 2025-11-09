package service

import (
	"basic-go/entity"
	"basic-go/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(ID string) (*entity.Category, error) {
	category := service.Repository.FindByID(ID)
	if category != nil {
		return category, nil
	}
	return nil, errors.New("category not found") 
}