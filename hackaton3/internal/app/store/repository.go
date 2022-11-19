package store

import "test_go/internal/app/model"

type Repository interface {
	Create(*model.User) error

	GetUsers() ([]model.User, error)

	FindUser(*model.User) (*model.User, error)
}
