package validatorutils_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
)

type ServiceSuite struct {
	suite.Suite
	validator *validator.Validate
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	validator, err := validatorutils.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}
	s.validator = validator.Validator

}

func (s *ServiceSuite) TestIsValidCpf() {
	cases := []struct {
		cpf      string
		expected bool
	}{
		{"11807116883", true},
		{"49694571820", true},
		{"68401726867", true},
		{"06035876820", true},
		{"00000000000", false},
		{"11111111111", false},
	}
	for _, test := range cases {
		err := s.validator.Var(test.cpf, "cpf")
		if test.expected {
			assert.Equal(s.T(), nil, err)
		} else {
			assert.NotEqual(s.T(), nil, err)
		}
	}
}
