package routes

import "github.com/micln/higo/context"

type Route struct {
	BaseRoute
	Handlers []Handler
	Method   []string
}

func (route *Route) Match(path string) bool {
	//path = strings.TrimRight(path, "/")
	return path == route.Pattern
}


func (route *Route) Go(ctx *context.Context, args ...string) {
	for _, handler := range route.Handlers {
		handler(ctx)
	}
}
