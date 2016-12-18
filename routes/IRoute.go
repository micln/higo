package routes

import "github.com/micln/higo/context"

//	A map from pattern to some handlers.
type IRoute interface {
	GetPattern() string
	Match(string) bool

	Go(*context.Context, ...string)
}

