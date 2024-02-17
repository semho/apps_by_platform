package handling

import (
	"fmt"
	"net/http"
	"platform/http/actionresults"
	"platform/services"
	"reflect"
	"regexp"
)

func (router *RouterComponent) AddMethodAlias(srcUrl string, method any, data ...any) *RouterComponent {
	var urlgen URLGenerator
	services.GetService(&urlgen)

	url, err := urlgen.GenerateUrl(method, data...)
	if err != nil {
		panic(err)
	}

	return router.AddUrlAlias(srcUrl, url)
}

func (router *RouterComponent) AddUrlAlias(srcUrl string, targetUrl string) *RouterComponent {
	aliasFunc := func(any) actionresults.ActionResult {
		return actionresults.NewRedirectAction(targetUrl)
	}
	alias := Route{
		httpMethod:  http.MethodGet,
		handlerName: "Alias",
		actionName:  "Redirect",
		expression:  *regexp.MustCompile(fmt.Sprintf("^%v[/]?$", srcUrl)),
		handlerMethod: reflect.Method{
			Type: reflect.TypeOf(aliasFunc),
			Func: reflect.ValueOf(aliasFunc),
		},
	}
	router.routes = append([]Route{alias}, router.routes...)
	return router
}
