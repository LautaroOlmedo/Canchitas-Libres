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

// Getters
func (f *Field) GetFieldID() string {
    return f.FieldID
}

func (f *Field) GetFieldType() string {
    return f.FieldType
}

func (f *Field) GetStatus() int {
    return f.Status
}

func (f *Field) GetFieldNumber() int {
    return f.FieldNumber
}

func (f *Field) GetFieldPrice() float64 {
    return f.FieldPrice
}

// Setters
func (f *Field) SetFieldID(fieldID string) {
    f.FieldID = fieldID
}

func (f *Field) SetFieldType(fieldType string) {
    f.FieldType = fieldType
}

func (f *Field) SetStatus(status int) {
    f.Status = status
}

func (f *Field) SetFieldNumber(fieldNumber int) {
    f.FieldNumber = fieldNumber
}

func (f *Field) SetFieldPrice(fieldPrice float64) {
    f.FieldPrice = fieldPrice
}