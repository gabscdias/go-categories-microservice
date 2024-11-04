package repositories

import (
	"github.com/gabscdias/go-categories-microservice/internal/entities"
)

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
