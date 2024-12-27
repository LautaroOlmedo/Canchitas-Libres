package field

import (
	"errors"
	"fmt"
	"testing"
)

func Test_NewField(t *testing.T) {
	type testCase struct {
		test 		string
		FieldType   string
		Status      int
		FieldNumber int
		FieldPrice  float64
		expectedError error
	}

	testCases := []testCase{
		{
		test: 		 "Prueba correcta. Ingresos de datos validos",     
		FieldType:   "Cancha 7",
		Status:      1,
		FieldNumber: 4,
		FieldPrice:  40000,
		expectedError: nil,
		},
		{
			test: 		 "Error: no ingreso el valor de la cancha",     
			FieldType:   "Cancha 7",
			Status:      1,
			FieldNumber: 4,
			FieldPrice:  0,
			expectedError: MissingParameterError,
		},
		{
			test: 		 "Error en el Estado de la Cancha",     
			FieldType:   "Cancha 7",
			Status:      0,
			FieldNumber: 4,
			FieldPrice:  40000,
			expectedError: MissingParameterError,
		},
		{
			test: 		 "Error no ingreso el tipo de cancha",     
			FieldType:   "",
			Status:      1,
			FieldNumber: 4,
			FieldPrice:  40000,
			expectedError: MissingParameterError,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			f, err := NewField(tc.FieldType, tc.Status, tc.FieldNumber, tc.FieldPrice)
			if !errors.Is(err, tc.expectedError) {
				fmt.Println("field: ", f)
				t.Fatalf("Yo esperaba el error: %v, y obtuve el error: %v", tc.expectedError, err)
			}
		})
	}
}