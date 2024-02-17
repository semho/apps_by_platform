package templates

import "io"

type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data any) (err error)
}
