package storage

import (
	domain "canchitas-libres/pkg/domain/field"
	"errors"
	"fmt"
)

func (m *Mongo) GetAll() ([]domain.Field, error) {
	fmt.Println("Getting all fields from mongodb")
	return []domain.Field{}, nil
}



func (m *Mongo) GetByFieldID(id int) (*domain.Field, error) {
	return nil, nil
}

func (m *Mongo) DeleteByFieldID(id int) error {
	return errors.New("Not implemented")
}