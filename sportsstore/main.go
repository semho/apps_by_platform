package main

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sportsstore/admin"
	"sportsstore/models/repo"
	"sportsstore/store"
	"sportsstore/store/cart"
	"sync"
)

func registerServices() {
	services.RegisterDefaultServices()
	//repo.RegisterMemoryRepoService()
	repo.RegisterSqlRepositoryService()
	sessions.RegisterSessionService()
	cart.RegisterCartService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},
		handling.NewRouter(
			handling.HandlerEntry{"", store.ProductHandler{}},
			handling.HandlerEntry{"", store.CategoryHandler{}},
			handling.HandlerEntry{"", store.CartHandler{}},
			handling.HandlerEntry{"", store.OrderHandler{}},
			handling.HandlerEntry{"admin", admin.AdminHandler{}},
			handling.HandlerEntry{"admin", admin.ProductsHandler{}},
			handling.HandlerEntry{"admin", admin.CategoriesHandler{}},
			handling.HandlerEntry{"admin", admin.OrdersHandler{}},
			handling.HandlerEntry{"admin", admin.DatabaseHandler{}},
		).AddMethodAlias("/", store.ProductHandler.GetProducts, 0, 1).
			AddMethodAlias("/products[/]?[A-z0-9]*?",
				store.ProductHandler.GetProducts, 0, 1).AddMethodAlias(
			"/admin[/]?", admin.AdminHandler.GetSection, ""),
	)
}

func main() {
	registerServices()
	result, err := services.Call(http.Serve, createPipeline())
	if err != nil {
		panic(err)
	}
	(result[0].(*sync.WaitGroup)).Wait()
}
