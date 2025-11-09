package repository

import (
	"basic-go/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindByID(ID string) *entity.Category {
	args := repository.Mock.Called(ID)

	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entity.Category)
	return &category
}
