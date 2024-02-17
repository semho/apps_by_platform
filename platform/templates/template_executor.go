package templates

import "io"

type InvokeHandlerFunc func(handlerName string, methodName string, args ...any) interface{}

type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data any) (err error)

	ExecTemplateWithFunc(writer io.Writer, name string, data any, handlerFunc InvokeHandlerFunc) (err error)
}
