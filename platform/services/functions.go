package services

import (
	"context"
	"errors"
	"reflect"
)

func Call(target any, otherArgs ...any) ([]any, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

func CallForContext(c context.Context, target any, otherArgs ...any) ([]any, error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Func {
		return nil, errors.New("Only functions can be invoked")
	}
	resultVals := invokeFunction(c, targetValue, otherArgs...)
	results := make([]any, len(resultVals))
	for i := 0; i < len(resultVals); i++ {
		results[i] = resultVals[i].Interface()
	}

	return results, nil
}
