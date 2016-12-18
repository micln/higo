package higo

import (
	"net/http"

	"github.com/micln/higo/builtins/middlewares"
	"github.com/micln/higo/context"
	"github.com/micln/higo/middleware"
	"github.com/micln/higo/routes"
)

type higo struct {
	Log ILogger
	middleware.MiddlewareLoader
	routes.RouteLoader
}

var DefaultServer = NewClassicHigo()

func NewClassicHigo() *higo {
	z := &higo{}
	z.Use(&middlewares.LogEntity{})
	return z
}

func (h *higo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := context.NewContext(w, r)
	defer func() {
		if r := recover(); r != nil {
			ctx.CatchException(r)
		}
	}()

	h.MiddlewareLoader.FireBefore(ctx)
	defer func() {
		h.MiddlewareLoader.FireAfter(ctx)
	}()

	found := false
	for _, v := range h.Routes {
		if v.Match(ctx.FullUrl()) {
			found = true
			v.Go(ctx)
			break
		}
	}

	if !found {
		ctx.ThrowHttpCode(404)
	}
}

func (h *higo) Run(addr string) {
	http.ListenAndServe(addr, h)
}
