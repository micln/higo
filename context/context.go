package context

import "net/http"

type Context struct {
	Req
	Resp
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{}
	ctx.Req.Request = r
	ctx.Resp.ResponseWriter = w
	ctx.Resp.SetContentType(`text/html`)
	ctx.Resp.SetStatusCode(`200`)
	ctx.Init()

	return ctx
}

func (ctx *Context) Init() {
}
