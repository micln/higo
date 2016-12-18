package routes

import (
	"fmt"
	"reflect"

	"github.com/micln/higo/context"
)

type AutoRoute struct {
	BaseRoute
	Funcs map[string]reflect.Value
}

func (this *AutoRoute) Match(path string) bool {
	return false
}

func (this *AutoRoute) Go(ctx *context.Context, args ...string) {
	httpMethod := ctx.Req.Request
	fmt.Println(httpMethod)
}
