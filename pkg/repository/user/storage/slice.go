package storage

import "canchitas-libres/pkg/domain/user"

func (s *Slice) GetAll() ([]user.User, error) {
	return []user.User{}, nil
}

func (s *Slice) GetByID(id int) (*user.User, error) {
	return nil, nil
}
