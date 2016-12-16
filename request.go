package higo

import (
	"net/http"
	"strings"
)

type Req struct {
	*http.Request
}

func (req *Req) FullUrl() string {
	return req.URL.Path
}

func (req *Req) Fragment() []string {
	return strings.Split(req.Request.URL.Path, "/")
}
