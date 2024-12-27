package storage

import (
	domain "canchitas-libres/pkg/domain/user"
	"fmt"
)

func (s *Slice) GetAll() ([]domain.User, error) {
	fmt.Println("Getting all users from slice")
	return []domain.User{}, nil
}

func (s *Slice) GetByID(id int) (*domain.User, error) {
	return nil, nil
}

func (S *Slice) DeleteByID(id int) error {
	return nil
}
