package zfgo

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

func (this *Route) Match(path string) bool {
	//path = strings.TrimRight(path, "/")
	return path == this.Pattern
}

type AutoRoute struct {
	BaseRoute
	Funcs map[string]reflect.Value
}

func (this *AutoRoute) Match(path string) bool {
	return false
}

func (this *Route) Go(ctx *Context, args ...string) {
	for _, handler := range this.Handlers {
		handler(ctx)
	}
}

func (this *AutoRoute) Go(ctx *Context, args ...string) {
	httpMethod := ctx.Req.Request
	fmt.Println(httpMethod)
}

func (this *zfgo) addRoute(path string, methods []string, handlers []Handler) {
	route := &Route{
		Method:   methods,
		Handlers: handlers,
	}
	route.Pattern = path
	this.Routes = append(this.Routes, route)
}

func (this *zfgo) addRouter(path string, c IController) {
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

	this.Routes = append(this.Routes, auto)
}

/*********	Public Functions	************/

func (this *zfgo) Get(pattern string, handlers ...Handler) {
	this.addRoute(pattern, []string{http.MethodGet}, handlers)
}

func (this *zfgo) Any(pattern string, handlers ...Handler) {
	this.addRoute(pattern, Methods, handlers)
}

func (this *zfgo) Router(pattern string, controller IController) {
	this.addRouter(pattern, controller)
}
