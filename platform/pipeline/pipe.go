package pipeline

import (
	"net/http"
	"platform/services"
	"reflect"
)

type RequestPipeline func(ctx *ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) {
	//pass
}

func CreatePipeline(components ...any) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		currentComponent := components[i]
		services.Populate(currentComponent)
		nextFunc := f
		if servComp, ok := currentComponent.(ServicesMiddlwareComponent); ok {
			f = createServiceDependentFunction(currentComponent, nextFunc)
			servComp.Init()
		} else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
			f = func(context *ComponentContext) {
				if context.error == nil {
					stdComp.ProcessRequest(context, nextFunc)
				}
			}
			stdComp.Init()
		} else {
			panic("Value is not a middleware component")
		}
	}
	return f
}

func createServiceDependentFunction(component any, nextFunc RequestPipeline) RequestPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if !method.IsValid() {
		panic("No ProcessRequestWithServices method defined")
	}

	return func(context *ComponentContext) {
		if context.error == nil {
			_, err := services.CallForContext(context.Request.Context(), method.Interface(), context, nextFunc)
			if err != nil {
				context.Error(err)
			}
		}
	}
}

func (pl RequestPipeline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
	deferredWriter := &DeferredResponseWriter{ResponseWriter: resp}
	ctx := ComponentContext{Request: req, ResponseWriter: deferredWriter}
	pl(&ctx)
	if ctx.error == nil {
		deferredWriter.FlushData()
	}
	return ctx.error
}
