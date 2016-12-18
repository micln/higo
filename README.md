# Higo

A go web framework like [laravel](http://laravel.com).

## Getting Start

```go
package main

import (
	utils "github.com/micln/go-utils"
	"github.com/micln/higo"
	"github.com/micln/higo/context"
)

var serv = higo.DefaultServer

func main() {
	serv.Get(`/`,
		func(ctx *context.Context) {
			ctx.WriteString(`It Works!`)
		},
		func(ctx *context.Context) {
			ctx.WriteString("<hr>")
			ctx.WriteString(utils.Date(`Y-m-d H:i:s`))
		},
	)

	serv.Run(`:8080`)
}

```

## Http Layer

- Routing
- Middleware
- Controller
- Request
- Response
- Context
- Session
- Token
- CSRF Protection
- External Tools
	- Validation
	- Cache
	- Redis
