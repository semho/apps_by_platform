package validation

import (
	"errors"
	"fmt"
	"strconv"
)

func required(fieldName string, value any, arg string) (bool, error) {
	str, ok := value.(string)
	if !ok {
		return false, errors.New("The required validator is for strings")
	}
	return str != "", fmt.Errorf("A value is required")
}

func min(fieldName string, value any, arg string) (valid bool, err error) {
	minVal, err := strconv.Atoi(arg)
	if err != nil {
		panic("Invalid arguments for validator: " + arg)
	}
	err = fmt.Errorf("The minimum value is %v", minVal)
	switch value.(type) {
	case int:
		valid = value.(int) >= minVal
	case float64:
		valid = value.(float64) >= float64(minVal)
	case string:
		err = fmt.Errorf("The minimal length is %v charachers", minVal)
		valid = len(value.(string)) >= minVal
	default:
		err = errors.New("The min validator is for int, float64, and str values")
	}

	return
}
