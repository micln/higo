package middleware

import "github.com/micln/higo/context"

type IMiddleware interface {
	Before(*context.Context)
	After(*context.Context)
}
