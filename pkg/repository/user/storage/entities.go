package storage

import (
	domain "canchitas-libres/pkg/domain/user"
	"time"
)

type Person struct {
	ID        string
	FirstName string
	LastName  string
	DNI       int
	BirthDate time.Time
}

type User struct {
	person   *Person
	email    string
	password string
	active   bool
	role     string
}

func (u *User) toDomain() *domain.User {
	var domainUser domain.User
	domainUser.SetEmail(u.email)
	domainUser.SetPassword(u.password)
	domainUser.SetActive(u.active)
	domainUser.SetRole(u.role)
	return &domainUser
}
