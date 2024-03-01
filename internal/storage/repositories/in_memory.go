package repositories

import (
	"fmt"
	"github.com/jollygrayskull/todo_list/internal/storage/models"
	"sort"
	"sync"
)

type IMRepository[T models.Entity] struct {
	data   map[int]T
	lastId int
	mutex  *sync.Mutex
}

func NewIMRepository[T models.Entity]() Repository[T] {
	return Repository[T](&IMRepository[T]{make(map[int]T), 1, &sync.Mutex{}})
}

func (repo *IMRepository[T]) Create(entity T) (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	id := repo.lastId
	entity.SetId(id)
	repo.data[id] = entity
	repo.lastId = id + 1

	return id, nil
}

func (repo *IMRepository[T]) Read(id int) (T, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	t, ok := repo.data[id]
	if !ok {
		var empty T
		return empty, fmt.Errorf("entry with id %d does not exist", id)
	}
	return t, nil
}

func (repo *IMRepository[T]) ReadAll() ([]T, error) {
	return repo.ReadAllSorted(func(left T, right T) bool {
		return left.GetId() < right.GetId()
	})
}

func (repo *IMRepository[T]) ReadAllSorted(compare func(T, T) bool) ([]T, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	var result []T

	for _, v := range repo.data {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})

	return result, nil
}

func (repo *IMRepository[T]) Count() (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return len(repo.data), nil
}

func (repo *IMRepository[T]) Update(entity T) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	id := entity.GetId()
	repo.data[id] = entity

	return nil
}

func (repo *IMRepository[T]) Delete(id int) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	delete(repo.data, id)

	return nil
}
