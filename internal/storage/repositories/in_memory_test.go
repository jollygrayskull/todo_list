package repositories

import (
	"github.com/jollygrayskull/todo_list/internal/storage/models"
	"testing"
)

type Person struct {
	models.BaseEntity
	FirstName string
}

func newPerson(firstName string) *Person {
	return &Person{
		FirstName: firstName,
	}
}

func TestIMRepository_Create(t *testing.T) {
	repo := NewIMRepository[*Person]()

	id, err := repo.Create(newPerson("John"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	p, err := repo.Read(id)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if p.FirstName != "John" {
		t.Errorf("first name is not John")
	}
	if id != 1 {
		t.Errorf("Returned id was not 1")
	}
	if p.Id != 1 {
		t.Errorf("id is not 1")
	}
}

func TestIMRepository_ReadAll(t *testing.T) {
	repo := NewIMRepository[*Person]()

	johnId, err := repo.Create(newPerson("John"))
	if err != nil {
		t.Errorf("error: %v", err)
	}
	bobId, err := repo.Create(newPerson("Bob"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	all, err := repo.ReadAll()
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if len(all) != 2 {
		t.Errorf("length is not 2")
	}
	if johnId != 1 {
		t.Errorf("john id should be 1")

	}
	if bobId != 2 {
		t.Errorf("bob id should be 2")
	}
}

func TestIMRepository_Count(t *testing.T) {
	repo := NewIMRepository[*Person]()

	_, err := repo.Create(newPerson("John"))
	if err != nil {
		t.Errorf("error: %v", err)
	}
	_, err = repo.Create(newPerson("Bob"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	count, err := repo.Count()
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if count != 2 {
		t.Errorf("count should be 2")
	}
}

func TestIMRepository_Update(t *testing.T) {
	repo := NewIMRepository[*Person]()

	id, err := repo.Create(newPerson("John"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	john, err := repo.Read(id)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	john.FirstName = "Steven"

	err = repo.Update(john)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	steven, err := repo.Read(id)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if steven.FirstName != "Steven" {
		t.Errorf("steven should be steven")
	}
}

func TestIMRepository_Delete(t *testing.T) {
	repo := NewIMRepository[*Person]()

	id, err := repo.Create(newPerson("John"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	err = repo.Delete(id)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	_, err = repo.Read(id)
	if err.Error() != "entry with id 1 does not exist" {
		t.Errorf("john was not deleted")
	}
}
