package validation

import (
	"reflect"
	"strings"
)

type TagValidator struct {
	validators map[string]ValidatorFunc
}

func (tv *TagValidator) Validate(data any) (ok bool, errs []ValidationError) {
	errs = []ValidationError{}
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Ptr {
		dataVal = dataVal.Elem()
	}
	if dataVal.Kind() != reflect.Struct {
		panic("Only struct can be validated")
	}
	for i := 0; i < dataVal.NumField(); i++ {
		fieldType := dataVal.Type().Field(i)
		validationTag, found := fieldType.Tag.Lookup("validation")
		if found {
			for _, v := range strings.Split(validationTag, ",") {
				var name, arg = "", ""
				if strings.Contains(v, ":") {
					nameAndArgs := strings.SplitN(v, ":", 2)
					name = nameAndArgs[0]
					arg = nameAndArgs[1]
				} else {
					name = v
				}

				if validator, ok := tv.validators[name]; ok {
					valid, err := validator(fieldType.Name, dataVal.Field(i).Interface(), arg)
					if !valid {
						errs = append(errs, ValidationError{
							FieldName: fieldType.Name,
							Error:     err,
						})
					}
				} else {
					panic("Unknown validator: " + name)
				}
			}
		}
	}
	ok = len(errs) == 0
	return
}

func NewDefaultValidator(validator map[string]ValidatorFunc) Validator {
	return &TagValidator{DefaultValidators()}
}
