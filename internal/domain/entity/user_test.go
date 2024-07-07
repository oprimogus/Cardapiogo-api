package entity_test

import (
	"testing"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

func TestIsValidUserRole(t *testing.T) {
	testCases := []struct {
		mockInput      string
		expectedOutput bool
		test           string
	}{
		{
			mockInput:      "consumer",
			expectedOutput: true,
			test:           "Test consumer role",
		},
		{
			mockInput:      "owner",
			expectedOutput: true,
			test:           "Test owner role",
		},
		{
			mockInput:      "employee",
			expectedOutput: true,
			test:           "Test employee role",
		},
		{
			mockInput:      "delivery_man",
			expectedOutput: true,
			test:           "Test delivery_man role",
		},
		{
			mockInput:      "admin",
			expectedOutput: true,
			test:           "Test admin role",
		},
		{
			mockInput:      "testeteste",
			expectedOutput: false,
			test:           "Test wrong role",
		},
	}

	for _, test := range testCases {
		result := entity.IsValidUserRole(test.mockInput)
		if result != test.expectedOutput {
			t.Errorf(`Error in case %s, expected %t but got %t`, test.test, test.expectedOutput, result)
		}
	}
}
