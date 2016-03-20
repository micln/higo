package zfgo

import (
	"net/http"
	"strings"

	"github.com/chanxuehong/wechat/json"
)

type Context struct {
	Req
	Resp
}

type Req struct {
	*http.Request
}

type Resp struct {
	http.ResponseWriter
}

func (resp *Resp) Writes(s string) {
	resp.Write([]byte(s))
}

func (this *Context) Init() {

}

func (this *Context) Fragment() []string {
	return strings.Split(this.Req.Request.URL.Path, "/")
}

func (this *Context) Error(code int, raw ...string) {
}

func (this *Context) Raw(raw string) {
	this.Resp.Writes(raw)
}

func (this *Context) Json(v interface{}) {
	s, _ := json.Marshal(v)
	this.Resp.Write(s)
}

func (this *Context) HTML(raw string) {
	this.Resp.Writes(raw)
}
