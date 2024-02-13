package services

import (
	"context"
	"errors"
	"reflect"
)

func GetService(target any) error {
	return GetServiceForContext(context.Background(), target)
}

func GetServiceForContext(c context.Context, target any) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr && targetValue.Elem().CanSet() {
		return resolveServiceFromValue(c, targetValue)
	}

	return errors.New("Type cannot be used as target")
}
