package reserve

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_NewReserve(t *testing.T) {
	type testCase struct {
		test          string
		userID        string
		fieldID       string
		date          time.Time
		reserveNumber int
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "Prueba correcta, ingreso todos los parametros.",
			userID:        "156123",
			fieldID:       "12213sdads",
			date:          time.Now(),
			reserveNumber: 456,
			expectedError: nil,
		},
		{
			test:          "Prueba incorrecta, no indico el ID del usuario.",
			userID:        "",
			fieldID:       "12213sdads",
			date:          time.Now(),
			reserveNumber: 456,
			expectedError: MissingParameterError,
		},
		{
			test:          "Prueba incorrecta, no indico el ID de la cancha.",
			userID:        "156123",
			fieldID:       "",
			date:          time.Now(),
			reserveNumber: 456,
			expectedError: MissingParameterError,
		},
		{
			test:          "Prueba incorrecta, falta el numero de reserva.",
			userID:        "156123",
			fieldID:       "12213sdads",
			date:          time.Now(),
			reserveNumber: 0,
			expectedError: MissingParameterError,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			r, err := NewReserve(tc.userID, tc.fieldID, tc.date, tc.reserveNumber)
			if !errors.Is(err, tc.expectedError) {
				fmt.Println("Reserve: ", r)
				t.Fatalf("Yo esperaba el error: %v, y obtuve el error: %v", tc.expectedError, err)
			}
		})
	}
}
