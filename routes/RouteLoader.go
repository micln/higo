package routes

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/micln/higo/controller"
)

type RouteLoader struct {
	Routes []IRoute
}

func (loader *RouteLoader) addRoute(path string, methods []string, handlers []Handler) {
	route := &Route{
		Method:   methods,
		Handlers: handlers,
	}
	route.Pattern = path
	loader.Routes = append(loader.Routes, route)
}

func (loader *RouteLoader) addRouter(path string, c controller.IController) {
	typ := reflect.TypeOf(c)

	auto := &AutoRoute{}
	auto.Pattern = path
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		for _, m := range Methods {
			if strings.Index(method.Name, m) == 0 {
				auto.Funcs[method.Name] = method.Func
			}
		}
	}

	loader.Routes = append(loader.Routes, auto)
}

/*********	Public Functions	************/

func (loader *RouteLoader) Get(pattern string, handlers ...Handler) {
	loader.addRoute(pattern, []string{http.MethodGet}, handlers)
}

func (loader *RouteLoader) Post(pattern string, handlers ...Handler) {
	loader.addRoute(pattern, []string{http.MethodPost}, handlers)
}

func (loader *RouteLoader) Any(pattern string, handlers ...Handler) {
	loader.addRoute(pattern, Methods, handlers)
}

func (loader *RouteLoader) Router(pattern string, controller controller.IController) {
	loader.addRouter(pattern, controller)
}
