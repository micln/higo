package routes

import "github.com/micln/higo/context"

type Handlers []Handler
type Handler func(*context.Context)
