package xvalidator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	xvalidator "github.com/oprimogus/cardapiogo/internal/api/validator"
)

type ServiceSuite struct {
	suite.Suite
	validator *xvalidator.Validator
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	validator, err := xvalidator.NewValidator("pt")
	if err != nil && validator == nil {
		panic(err)
	}
	s.validator = validator
}

func (s *ServiceSuite) TestIsValidCpf() {
	transactionID := "transactionTest"
	cases := []struct {
		Cpf      string `validate:"cpf"`
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
		err := s.validator.Validate(test, transactionID)
		if test.expected {
			assert.Nil(s.T(), err, fmt.Sprintf("Expect nil and got %s for cpf %s", err, test.Cpf))
		} else {
			assert.NotNil(s.T(), err, fmt.Sprintf("Expect error and got %s for cpf %s", err, test.Cpf))
		}
	}
}

func (s *ServiceSuite) TestIsValidCnpj() {
	transactionID := "transactionTest"
	cases := []struct {
		Cnpj     string `validate:"cnpj"`
		Expected bool
	}{
		{"46277350000109", true},
		{"14380200000121", true},
		{"11111111111111", false},
		{"00000000000000", false},
		{"134676845867423", false},
		{"243661654000166", false},
	}

	for _, test := range cases {
		err := s.validator.Validate(test, transactionID)
		if test.Expected {
			assert.Nil(s.T(), err, fmt.Sprintf("Expect nil and got %s for CNPJ %s", err, test.Cnpj))
		} else {
			assert.NotNil(s.T(), err, fmt.Sprintf("Expect error and got %s for CNPJ %s", err.ErrorMessage, test.Cnpj))
		}
	}
}

func (s *ServiceSuite) TestIsValidPhone() {
	transactionID := "transactionTest"
	cases := []struct {
		Phone    string `validate:"required,phone"`
		expected bool
	}{
		{"+5513981488616", true},
		{"+banana", false},
		{"-5501123456789", false},
		{"+551381486097534", false},
		{"1234567890", false},
	}
	for _, test := range cases {
		err := s.validator.Validate(test, transactionID)
		if test.expected {
			assert.Nil(s.T(), err, fmt.Sprintf("Expect nil and got %s for phone %s", err, test.Phone))
		} else {
			assert.NotNil(s.T(), err, fmt.Sprintf("Expect error and got %s for phone %s", err, test.Phone))
		}
	}
}
