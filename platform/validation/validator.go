package validation

type ValidationError struct {
	FieldName string
	Error     error
}

type Validator interface {
	Validate(data any) (ok bool, err []ValidationError)
}

type ValidatorFunc func(fieldName string, value any, arg string) (bool, error)

func DefaultValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required": required,
		"min":      min,
	}
}
