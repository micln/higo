package higo

import "net/http"

type Context struct {
	Req
	Resp
}

func (ctx *Context) Init(w http.ResponseWriter, r *http.Request) {
	ctx.Req.Request = r
	ctx.Resp.ResponseWriter = w
}

