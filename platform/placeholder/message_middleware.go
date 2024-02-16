package placeholder

import (
	"platform/config"
	"platform/pipeline"
	"platform/services"
	"platform/templates"
)

type SimpleMessageComponent struct {
	Message string
	config.Configuration
}

func (c *SimpleMessageComponent) Init() {
	services.GetService(&c.Configuration)
	c.Message = c.Configuration.GetStringDefault("main:message", "Default Message")
}

func (c *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (c *SimpleMessageComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	executor templates.TemplateExecutor) {
	err := executor.ExecTemplate(ctx.ResponseWriter,
		"simple_message.html", c.Message)
	if err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
