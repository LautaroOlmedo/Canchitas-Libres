package reserve

import (
	"errors"
	"time"

	"github.com/google/uuid"
	// "canchitas-libres/pkg/domain/field"
	// "canchitas-libres/pkg/domain/user"
)

var (
	MissingParameterError = errors.New("missing parameter")
)

type Reserve struct {
	reserveID     string
	userID        string
	fieldID       string
	date          time.Time
	reserveNumber int
}

// Nueva reserva
func NewReserve(userID string, fieldID string, date time.Time, reserveNumber int) (*Reserve, error) {

	if userID == "" || fieldID == "" {
		return nil, MissingParameterError
	}
	if reserveNumber == 0 {
		return nil, MissingParameterError
	}
	reserveUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Reserve{
		reserveID:     reserveUUID.String(),
		userID:        userID,  //Se trabaja como un string pero a la hora de llamar a la funcion se envia el parametro (*user.User).GetUserID(nil)???
		fieldID:       fieldID, // mismo caso que con UserID
		date:          date,
		reserveNumber: reserveNumber,
	}, nil
}
