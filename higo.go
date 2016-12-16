package higo

import (
	"encoding/json"
	"fmt"
	`log`
	"net/http"
)

type higo struct {
	Routes []IRoute
}

var DefaultServer = NewHigo()

func NewHigo() *higo {
	z := &higo{}
	return z
}

func (h *higo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	s, _ := json.Marshal(r.URL)

	fmt.Printf("%s\n", s)

	ctx := &Context{
	}
	ctx.Init(w, r)

	defer func() {
		ctx.CatchException()
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
		ctx.ThrowHttp(404)
	}
}

func (h *higo) Run(addr string) {
	http.ListenAndServe(addr, h)
}
