package user

import (
	"canchitas-libres/pkg/user/domain/person"
	"time"
)

type User struct {
	person   *person.Person
	email    string
	password string
	active   bool
	role     string
}

func NewUser(firstName string, lastName string, DNI int, birthDate time.Time, email, password string, role string) User {
	person := &person.Person{
		FirstName: firstName,
		LastName:  lastName,
		DNI:       DNI,
		BirthDate: birthDate,
	}
	return User{
		person:   person,
		email:    email,
		password: password,
		active:   true,
		role:     role,
	}
}

func (u *User) GetPerson() *person.Person {
	return u.person
}

func (u *User) GetEmail() string {
	return u.email
}
