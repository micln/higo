package higo

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func (resp *Resp) SetType(typ string){
	resp.Header().Set(`Content-Type`,typ)
}

func (resp *Resp) WriteString(s string) {
	resp.Write([]byte(s))
}

func (ctx *Context) Init() {

}

func (ctx *Context) Fragment() []string {
	return strings.Split(ctx.Req.Request.URL.Path, "/")
}

func (ctx *Context) Error(code int, raw ...string) {
}

//	@todo 用XSS过滤
func (ctx *Context) Raw(raw string) {
	ctx.Resp.SetType(`text/plain`)
	ctx.Resp.WriteString(raw)
}

func (ctx *Context) JSON(vs ...interface{}) {
	var s []byte
	if len(vs) > 1 {
		s, _ = json.Marshal(vs)
	}else{
		s, _ = json.Marshal(vs[0])
	}
	ctx.Resp.Write(s)
}

func (ctx *Context) XML(v interface{}) {
	s, _ := xml.Marshal(v)
	ctx.Resp.Write(s)
}

func (ctx *Context) HTML(raw string) {
	ctx.Resp.WriteString(raw)
}

func (ctx * Context) File(filename string){
	f,err:=os.Open(filename)
	assert(err)

	//fi,err:=f.Stat()
	//assert(err)

	b,err:=ioutil.ReadAll(f)
	assert(err)

	ctx.Resp.Write(b)
}
