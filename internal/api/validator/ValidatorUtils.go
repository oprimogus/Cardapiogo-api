package validatorutils

import (
	"fmt"
	"net/http"

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
	validator  *validator.Validate
	translator ut.Translator
	locale     string
}

func NewValidator(locale string) (*Validator, error) {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterValidation("role", isValidUserRole)
	v.RegisterValidation("account_provider", isValidAccountProvider)

	enLocale := en.New()
	ptLocale := pt.New()
	uni := ut.New(enLocale, ptLocale, enLocale)

	translator, found := uni.GetTranslator(locale)
	if !found {
		return nil, fmt.Errorf("Locale %s not found", locale)
	}
	switch locale {
	case "en":
		en_translations.RegisterDefaultTranslations(v, translator)
	case "pt":
		pt_translations.RegisterDefaultTranslations(v, translator)
	default:
		return nil, fmt.Errorf("unsupported locale: %s", locale)
	}

	return &Validator{
		validator:  v,
		translator: translator,
		locale:     locale,
	}, nil
}

func (v *Validator) Validate(i interface{}) *errors.ErrorResponse {
	out := make(map[string]string)

	// Realiza a validação
	err := v.validator.Struct(i)

	// Verifica se houve erros de validação
	if err != nil {
		// Realiza o type assertion apenas se houver erros
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// Se a asserção de tipo falhar, retorna um erro genérico ou lida com isso de maneira apropriada
			out["error"] = "Unknown validation error"
			return errors.NewErrorResponse(http.StatusBadRequest, out["error"])
		}

		// Processa os erros de validação
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
		return "Valor inválido para o campo."
	}
	return "Invalid value for field"
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
	role := fl.Field().String()
	switch types.AccountProvider(role) {
	case types.AccountProviderApple,
		types.AccountProviderGoogle,
		types.AccountProviderMeta:
		return true
	default:
		return false
	}
}
