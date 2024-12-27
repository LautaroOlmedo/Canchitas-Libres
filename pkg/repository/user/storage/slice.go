package storage

import (
	domain "canchitas-libres/pkg/domain/user"
)

func (s *Slice) GetAll() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (s *Slice) GetByID(id int) (*domain.User, error) {
	return nil, nil
}
