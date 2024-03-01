package repositories

import (
	"github.com/jollygrayskull/todo_list/internal/storage/models"
)

type Repository[T models.Entity] interface {
	Create(entity T) (int, error)
	Read(id int) (T, error)
	ReadAll() ([]T, error)
	ReadAllSorted(compare func(T, T) bool) ([]T, error)
	Count() (int, error)
	Update(entity T) error
	Delete(id int) error
}
