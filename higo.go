package higo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type higo struct {
	Routes []IRoute
}

func NewHigo() *higo {
	z := &higo{}
	return z
}

func (h *higo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	s, _ := json.Marshal(r.URL)

	fmt.Printf("%s\n", s)

	req := &Context{}
	req.Req.Request = r
	req.Resp.ResponseWriter = w
	req.Init()

	path := r.URL.Path

	found := false
	for _, v := range h.Routes {
		if v.Match(path) {
			found = true
			v.Go(req)
			break
		}
	}

	if !found {
		req.Raw(`404`)
	}
}



func (h *higo) Run(addr string) {
	http.ListenAndServe(addr, h)
}
