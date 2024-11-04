package usecases

import (
	"github.com/gabscdias/go-categories-microservice/internal/entities"
	"github.com/gabscdias/go-categories-microservice/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
