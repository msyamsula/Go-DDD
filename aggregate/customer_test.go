package aggregate_test

import (
	"testing"

	"github.com/msyamsula/Go-DDD/aggregate"
)

func Test_NewCustomer(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "valid Person",
			name:        "syamsul",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if err != tc.expectedErr {
				t.Errorf("Expected: %v, Got: %v", tc.expectedErr, err)
			}
		})
	}

}
