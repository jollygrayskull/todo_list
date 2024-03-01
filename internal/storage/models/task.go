package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Priority int

const (
	High Priority = iota
	Medium
	Low
)

func priorityFromString(str string) (Priority, bool) {
	priority, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return -1, false
	}

	return priorityFromInt(priority)
}

func priorityFromInt(priority int64) (Priority, bool) {
	switch priority {
	case 0:
		return High, true
	case 1:
		return Medium, true
	case 2:
		return Low, true
	default:
		return -1, false
	}
}

func (p Priority) String() string {
	switch p {
	case 0:
		return "High"
	case 1:
		return "Medium"
	case 2:
		return "Low"
	default:
		return "Invalid Priority"
	}
}

type Category string

const (
	Work      Category = "Work"
	Household Category = "Household"
	Personal  Category = "Personal"
)

func categoryFromString(str string) (Category, bool) {
	switch {
	case strings.EqualFold(str, string(Work)):
		return Work, true
	case strings.EqualFold(str, string(Household)):
		return Household, true
	case strings.EqualFold(str, string(Personal)):
		return Personal, true
	default:
		return "", false
	}
}

type Task struct {
	BaseEntity
	Title       string
	Description string
	Priority    Priority
	Category    Category
	Completed   bool
}

func NewTask(title string, description string, priority string, category string) (Task, error) {
	prior, ok := priorityFromString(priority)
	if !ok {
		return Task{}, fmt.Errorf("invalid priority %s", priority)
	}
	cat, ok := categoryFromString(category)
	if !ok {
		return Task{}, fmt.Errorf("invalid category %s", category)
	}

	task := Task{
		Title:       title,
		Description: description,
		Priority:    prior,
		Category:    cat,
		Completed:   false,
	}

	return task, nil
}

func (t Task) String() string {
	return fmt.Sprintf("id: %d, title: %s, description: %s, priority: %d, category %s",
		t.Id, t.Title, t.Description, t.Priority, t.Category)
}
