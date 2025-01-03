package field

import (
	"errors"

	"github.com/google/uuid"
)

type Field struct {
	FieldID     string
	FieldType   string
	Status      int
	FieldNumber int
	FieldPrice  float64
}

var (
	MissingParameterError = errors.New("missing parameter")
)

// crear patron factory para Field
func NewField(fieldType string, status int, fieldNumber int, fieldPrice float64) (*Field, error) {

	if fieldType == "" {
		return nil, MissingParameterError
	}

	if fieldPrice == 0 || fieldNumber == 0 {
		return nil, MissingParameterError
	}

	if status > 3 || status <= 0 {
		return nil, MissingParameterError
	}

	fieldID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Field{
		FieldID:     fieldID.String(),
		FieldType:   fieldType,
		Status:      status,
		FieldNumber: fieldNumber,
		FieldPrice:  fieldPrice,
	}, nil
}