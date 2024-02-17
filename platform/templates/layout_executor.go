package templates

import (
	"html/template"
	"io"
	"strings"
)

type LayoutTemplateProcessor struct{}

var emptyFunc = func(handlerName, methodName string, args ...any) any { return "" }

func (proc *LayoutTemplateProcessor) ExecTemplateWithFunc(writer io.Writer,
	name string, data any, handlerFunc InvokeHandlerFunc) (err error) {

	var sb strings.Builder
	layoutName := ""
	localTemplates := getTemplates()
	localTemplates.Funcs(map[string]any{
		"body":    insertBodyWrapper(&sb),
		"layout":  setLayoutWrapper(&layoutName),
		"handler": handlerFunc,
	})
	err = localTemplates.ExecuteTemplate(&sb, name, data)
	if layoutName != "" {
		localTemplates.ExecuteTemplate(writer, layoutName, data)
	} else {
		io.WriteString(writer, sb.String())
	}
	return
}

func (proc *LayoutTemplateProcessor) ExecTemplate(writer io.Writer,
	name string, data any) (err error) {
	return proc.ExecTemplateWithFunc(writer, name, data, emptyFunc)
}

var getTemplates func() (t *template.Template)

func insertBodyWrapper(body *strings.Builder) func() template.HTML {
	return func() template.HTML {
		return template.HTML(body.String())
	}
}
func setLayoutWrapper(val *string) func(string) string {
	return func(layout string) string {
		*val = layout
		return ""
	}
}
