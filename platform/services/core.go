package services

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

func addService(life lifecycle, factoryFunc any) error {

	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() != reflect.Func || factoryFuncType.NumOut() != 1 {
		return fmt.Errorf("Type cannot be used as service: %v", factoryFuncType)
	}
	services[factoryFuncType.Out(0)] = BindingMap{
		factoryFunc: reflect.ValueOf(factoryFunc),
		lifecycle:   life,
	}
	return nil
}

var contextReference = (*context.Context)(nil)
var contextReferenceType = reflect.TypeOf(contextReference).Elem()

func resolveServiceFromValue(c context.Context, val reflect.Value) error {
	var err error
	serviceType := val.Elem().Type()
	if serviceType == contextReferenceType {
		val.Elem().Set(reflect.ValueOf(c))
	} else if binding, found := services[serviceType]; found {
		if binding.lifecycle == Scoped {
			resolveScopedService(c, val, binding)
		} else {
			val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
		}
	} else {
		err = fmt.Errorf("Cannot find service %v", serviceType)
	}

	return err
}

func resolveScopedService(c context.Context, val reflect.Value, binding BindingMap) error {
	sMap, ok := c.Value(ServiceKey).(serviceMap)
	if !ok {
		val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
		return nil
	}
	serviceVal, ok := sMap[val.Type()]
	if !ok {
		serviceVal = invokeFunction(c, binding.factoryFunc)[0]
		sMap[val.Type()] = serviceVal
	}
	val.Elem().Set(serviceVal)

	return nil
}

func resolveFunctionArguments(c context.Context, f reflect.Value, otherArgs ...any) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	if otherArgs != nil {
		for ; i < len(otherArgs); i++ {
			params[i] = reflect.ValueOf(otherArgs[i])
		}
	}

	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		err := resolveServiceFromValue(c, pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}

	return params
}

func invokeFunction(c context.Context, f reflect.Value, otherArgs ...any) []reflect.Value {
	return f.Call(resolveFunctionArguments(c, f, otherArgs...))
}
