package person

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	MissingParameterError = errors.New("missing parameter")
)

type Person struct {
	ID        string
	FirstName string
	LastName  string
	DNI       int
	BirthDate time.Time
}

// patron Factory: Es una fabrica que rotorna el tipo de entidad
func NewPerson(firstName string, lastName string, DNI int, birthDate time.Time) (*Person, error) {
	if firstName == "" || lastName == "" {
		return nil, MissingParameterError
	}
	if birthDate.IsZero() {
		return nil, MissingParameterError
	}

	if DNI == 0 {
		return nil, MissingParameterError
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Person{
		ID:        id.String(),
		FirstName: firstName,
		LastName:  lastName,
		DNI:       DNI,
		BirthDate: birthDate,
	}, nil
}
