package todo

import (
	"errors"
	"time"
	"todoapp/pkg/todoapp"
)

var ErrInvalidTask = errors.New("invalid task")

type Service struct {
	db DB
}

func NewService(db DB) *Service {
	return &Service{db: db}
}

type DB interface {
	Insert(task *todoapp.Task) error
	All() ([]*todoapp.Task, error)
}

func (s *Service) Insert(task *todoapp.Task) error {
	if task.Value == "" {
		return ErrInvalidTask
	}
	return s.db.Insert(task)
}

func (s *Service) FutureTasks() ([]*todoapp.Task, error) {
	tasks, err := s.db.All()
	if err != nil {
		return nil, err
	}

	var result []*todoapp.Task
	for _, task := range tasks {
		if task.Date.After(time.Now()) {
			result = append(result, task)
		}
	}
	return result, nil
}
