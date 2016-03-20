package zfgo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/json"
)

type zfgo struct {
	Routes []IRoute
}

func New() *zfgo {
	z := &zfgo{}
	return z
}

func (this *zfgo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	s, _ := json.Marshal(r.URL)

	fmt.Printf("%s\n", s)

	req := &Context{}
	req.Req.Request = r
	req.Resp.ResponseWriter = w
	req.Init()

	path := r.URL.Path

	found := false
	for _, v := range this.Routes {
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

func (this *zfgo) Run(addr string) {
	http.ListenAndServe(addr, this)
}
