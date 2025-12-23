package repository

import "todo-app/internal/entity"

type TaskRepository interface {
	Create(task *entity.Task) error
	GetAll() ([]*entity.Task, error)
}

type LocalStorage struct {
	db []*entity.Task
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		db: make([]*entity.Task, 0),
	}
}

func (l *LocalStorage) Create(task *entity.Task) error {
	l.db = append(l.db, task)
	return nil
}

func (l *LocalStorage) GetAll() ([]*entity.Task, error) {
	return l.db, nil
}
