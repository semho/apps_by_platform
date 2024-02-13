package services

import (
	"context"
	"errors"
	"reflect"
)

func Populate(target any) error {
	return PopulateForContext(context.Background(), target)
}

func PopulateForContext(c context.Context, target any) error {
	return PopulateForContextWithExtras(c, target, make(map[reflect.Type]reflect.Value))
}

func PopulateForContextWithExtras(c context.Context, target any, extras map[reflect.Type]reflect.Value) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Struct {
		return errors.New("Type cannot be used as target")
	}
	targetValue = targetValue.Elem()
	for i := 0; i < targetValue.Type().NumField(); i++ {
		fieldVal := targetValue.Field(i)
		if !fieldVal.CanSet() {
			continue
		}
		if extra, ok := extras[fieldVal.Type()]; ok {
			fieldVal.Set(extra)
		} else {
			err := resolveServiceFromValue(c, fieldVal.Addr())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
