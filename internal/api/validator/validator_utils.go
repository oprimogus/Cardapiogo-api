package validatorutils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	pt_translations "github.com/go-playground/validator/v10/translations/pt"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

var personalizedValidations = map[string]bool{
	"role":             true,
	"account_provider": true,
}

type Validator struct {
	Validator  *validator.Validate
	translator ut.Translator
	locale     string
}

func NewValidator(locale string) (*Validator, error) {
	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.RegisterValidation("role", isValidUserRole)
	if err != nil {
		panic(fmt.Sprintf("Could not create validator for type role: %v", err))
	}
	err = v.RegisterValidation("account_provider", isValidAccountProvider)
	if err != nil {
		panic(fmt.Sprintf("Could not create validator for type account_provider: %v", err))
	}
	err = v.RegisterValidation("cpf", IsValidCpf)
	if err != nil {
		panic(fmt.Sprintf("Could not create validator for type cpf: %v", err))
	}

	enLocale := en.New()
	ptLocale := pt.New()
	uni := ut.New(enLocale, ptLocale, enLocale)

	translator, found := uni.GetTranslator(locale)
	if !found {
		return nil, fmt.Errorf("locale %s not found", locale)
	}
	switch locale {
	case "en":
		err = en_translations.RegisterDefaultTranslations(v, translator)
		if err != nil {
			panic(fmt.Sprintf("Could not register locale %v translation: %v", locale, err))
		}
	case "pt":
		err = pt_translations.RegisterDefaultTranslations(v, translator)
		if err != nil {
			panic(fmt.Sprintf("Could not register locale %v translation: %v", locale, err))
		}
	default:
		return nil, fmt.Errorf("unsupported locale: %s", locale)
	}

	return &Validator{
		Validator:  v,
		translator: translator,
		locale:     locale,
	}, nil
}

func (v *Validator) Validate(i interface{}) *errors.ErrorResponse {
	out := make(map[string]string)

	err := v.Validator.Struct(i)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			out["error"] = "Unknown validation error"
			return errors.New(http.StatusBadRequest, out["error"])
		}

		for _, e := range errs {
			_, isPersonalized := personalizedValidations[e.Tag()]
			if isPersonalized {
				out[e.StructField()] = errorPersonalized(v.locale, e.Tag())
			} else {
				out[e.StructField()] = e.Translate(v.translator)
			}
		}
	}

	if len(out) > 0 {
		return errors.InvalidInput(out)
	}
	return nil
}

func errorPersonalized(locale string, tag string) string {
	if locale == "pt" {
		return fmt.Sprintf("Valor inválido para o campo %v.", tag)
	}
	return fmt.Sprintf("Invalid value for field %v", tag)
}

func isValidUserRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	switch types.Role(role) {
	case types.UserRoleAdmin,
		types.UserRoleConsumer,
		types.UserRoleDeliveryMan,
		types.UserRoleEmployee,
		types.UserRoleOwner:
		return true
	default:
		return false
	}
}

func isValidAccountProvider(fl validator.FieldLevel) bool {
	accountProvider := fl.Field().String()
	switch types.AccountProvider(accountProvider) {
	case types.AccountProviderApple,
		types.AccountProviderGoogle,
		types.AccountProviderMeta:
		return true
	default:
		return false
	}
}

func IsValidCpf(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()

	if len(cpf) != 11 {
		return false
	}
	if isAllEqual(cpf) {
		return false
	}
	d1 := calculateDigitCpf(cpf, 10)
	d2 := calculateDigitCpf(cpf, 11)
	return strconv.Itoa(d1) == cpf[9:10] && strconv.Itoa(d2) == cpf[10:11]
}

func isAllEqual(value string) bool {
	for i := range value {
		if value[i] != value[0] {
			return false
		}
	}
	return true
}

func calculateDigitCpf(cpf string, weight int) int {
	sum := 0
	count := weight - 1
	for i := 0; i < count; i++ {
		number, _ := strconv.Atoi(string(cpf[i]))
		sum += number * weight
		weight--
	}
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func calculateDigitCnpj(cnpj string, factor int) int {
	sum := 0
	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := 0; i < factor-1; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weights[i+12-factor]
	}
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func IsValidCnpj(fl validator.FieldLevel) bool {
	cnpj := fl.Field().String()

	if len(cnpj) != 14 {
		return false
	}
	if isAllEqual(cnpj) {
		return false
	}
	d1 := calculateDigitCnpj(cnpj, 13)
	d2 := calculateDigitCnpj(cnpj, 14)
	return strconv.Itoa(d1) == cnpj[12:13] && strconv.Itoa(d2) == cnpj[13:14]
}

func IsValidCpfOrCnpj(fl validator.FieldLevel) bool {
	return IsValidCpf(fl) || IsValidCnpj(fl)
}
