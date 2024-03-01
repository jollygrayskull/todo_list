package storage

import (
	"github.com/jollygrayskull/todo_list/internal/storage/models"
	"github.com/jollygrayskull/todo_list/internal/storage/repositories"
)

type Storage struct {
	TaskRepo repositories.Repository[*models.Task]
}

func NewIMStorage() *Storage {
	return &Storage{
		TaskRepo: repositories.NewIMRepository[*models.Task](),
	}
}
