package storage

import (
	domain "canchitas-libres/pkg/domain/field"
)

type Field struct {
	FieldID     string
	FieldType   string
	Status      int
	FieldNumber int
	FieldPrice  float64
}


func (f *Field) toDomain() *domain.Field {
	var domainField domain.Field
	domainField.SetFieldID(f.FieldID)
	domainField.SetFieldType(f.FieldType)
	domainField.SetFieldNumber(f.FieldNumber)
	domainField.SetFieldPrice(f.FieldPrice)
	domainField.SetStatus(f.Status)
	return &domainField
}