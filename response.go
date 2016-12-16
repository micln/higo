package higo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/micln/higo/exception"
)

type Resp struct {
	http.ResponseWriter
}

func (resp *Resp) SetType(typ string) {
	resp.Header().Set(`Content-Type`, typ)
}

func (resp *Resp) WriteString(s string) {
	resp.Write([]byte(s))
}

//	@todo 用XSS过滤
func (resp *Resp) Raw(raw string) {
	resp.SetType(`text/plain`)
	resp.WriteString(raw)
}

func (resp *Resp) JSON(vs ...interface{}) {
	var s []byte
	if len(vs) > 1 {
		s, _ = json.Marshal(vs)
	} else {
		s, _ = json.Marshal(vs[0])
	}
	resp.Write(s)
}

func (resp *Resp) XML(v interface{}) {
	s, _ := xml.Marshal(v)
	resp.Write(s)
}

func (resp *Resp) HTML(raw string) {
	resp.WriteString(raw)
}

func (resp *Resp) File(filename string) {
	f, err := os.Open(filename)
	resp.AssertNil(err)

	b, err := ioutil.ReadAll(f)
	resp.AssertNil(err)

	resp.Write(b)
}

func (resp *Resp) ThrowException(e exception.IException) {
	panic(e)
}

func (resp *Resp) Throw(code int, message string) {
	resp.ThrowException(&exception.DefaultException{
		Code:    code,
		Message: message,
	})
}

var httpCodeMessage = map[int]string{
	404:`Not Found`,
}

func (resp *Resp) ThrowHttp(code int) {
	msg := httpCodeMessage[code]
	if len(msg) < 0 {
		resp.Throw(500, fmt.Sprintf("Unknow http code [%s]", code))
	}
	resp.Throw(code, httpCodeMessage[code])
}

func (resp *Resp) AssertNil(err error) {
	if err != nil {
		resp.Throw(500, err.Error())
	}
}

func (resp *Resp) CatchException() {
	if r := recover(); r != nil {
		resp.ResponseException(r)
	}
}

func (resp *Resp) ResponseException(r interface{}) {
	exc, ok := r.(exception.IException)

	if !ok {
		exc = &exception.DefaultException{
			Code:    500,
			Message: fmt.Sprint(r),
		}
	}

	resp.WriteString(exc.Response())
}
