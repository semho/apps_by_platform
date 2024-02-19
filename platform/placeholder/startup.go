package placeholder

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},
		//&SimpleMessageComponent{},
		handling.NewRouter(
			handling.HandlerEntry{"", NameHandler{}},
			handling.HandlerEntry{"", DayHandler{}},
		).AddMethodAlias("/", NameHandler.GetNames),
	)
}

func Start() {
	sessions.RegisterSessionService()
	results, err := services.Call(http.Serve, createPipeline())
	if err != nil {
		panic(err)
	}
	results[0].(*sync.WaitGroup).Wait()
}
