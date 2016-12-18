package middleware

import "github.com/micln/higo/context"


type MiddlewareLoader struct {
	middlewares []IMiddleware

	hasFireBefore bool
	hasFireAfter bool
}

func (loader *MiddlewareLoader) Use(mid IMiddleware){
	loader.middlewares = append(loader.middlewares, mid)
}

func (loader *MiddlewareLoader) FireBefore(ctx *context.Context){
	for _, mid := range loader.middlewares {
		mid.Before(ctx)
	}
	loader.hasFireBefore = true
}

func (loader *MiddlewareLoader) FireAfter(ctx *context.Context){
	for _, mid := range loader.middlewares {
		mid.After(ctx)
	}
	loader.hasFireAfter = true
}
