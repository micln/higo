package higo

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var Methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
}

type IRoute interface {
	GetPattern() string
	Match(string) bool

	Go(*Context, ...string)
}

type BaseRoute struct {
	Pattern string
}

func (this *BaseRoute) GetPattern() string {
	return this.Pattern
}

type Route struct {
	BaseRoute
	Handlers []Handler
	Method   []string
}

func (route *Route) Match(path string) bool {
	//path = strings.TrimRight(path, "/")
	return path == route.Pattern
}

type AutoRoute struct {
	BaseRoute
	Funcs map[string]reflect.Value
}

func (this *AutoRoute) Match(path string) bool {
	return false
}

func (route *Route) Go(ctx *Context, args ...string) {
	for _, handler := range route.Handlers {
		handler(ctx)
	}
}

func (this *AutoRoute) Go(ctx *Context, args ...string) {
	httpMethod := ctx.Req.Request
	fmt.Println(httpMethod)
}

func (h *higo) addRoute(path string, methods []string, handlers []Handler) {
	route := &Route{
		Method:   methods,
		Handlers: handlers,
	}
	route.Pattern = path
	h.Routes = append(h.Routes, route)
}

func (h *higo) addRouter(path string, c IController) {
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

	h.Routes = append(h.Routes, auto)
}

/*********	Public Functions	************/

func (h *higo) Get(pattern string, handlers ...Handler) {
	h.addRoute(pattern, []string{http.MethodGet}, handlers)
}

func (h *higo) Post(pattern string, handlers ...Handler) {
	h.addRoute(pattern, []string{http.MethodPost}, handlers)
}

func (h *higo) Any(pattern string, handlers ...Handler) {
	h.addRoute(pattern, Methods, handlers)
}

func (h *higo) Router(pattern string, controller IController) {
	h.addRouter(pattern, controller)
}
