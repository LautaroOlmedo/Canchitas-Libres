package storage

import (
	domain "canchitas-libres/pkg/domain/user"
	"errors"
	"fmt"
)

func (m *Mongo) GetAll() ([]domain.User, error) {
	fmt.Println("Getting all users from mongo")
	return []domain.User{}, nil
}

func (m *Mongo) GetByID(id int) (*domain.User, error) {
	return nil, nil
}

func (m *Mongo) DeleteByID(id int) error {
	return errors.New("Not implemented")
}


