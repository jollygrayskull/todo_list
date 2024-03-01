package models

type Entity interface {
	GetId() int
	SetId(id int)
}

type BaseEntity struct {
	Id int
}

func (t *BaseEntity) GetId() int {
	return t.Id
}

func (t *BaseEntity) SetId(id int) {
	t.Id = id
}
