package user

import (
	"canchitas-libres/pkg/domain/person"
	"time"
)

type User struct {
	person   *person.Person
	email    string
	password string
	active   bool
	role     string
}

// NewUser creates a new User instance
func NewUser(firstName string, lastName string, DNI int, birthDate time.Time, email, password string, role string) (User, error) {
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
	}, nil
}

// Getters
func (u *User) GetPerson() *person.Person {
	return u.person
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) IsActive() bool {
	return u.active
}

func (u *User) GetRole() string {
	return u.role
}

// Setters
func (u *User) SetPerson(person *person.Person) {
	u.person = person
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetActive(active bool) {
	u.active = active
}

func (u *User) SetRole(role string) {
	u.role = role
}
