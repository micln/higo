package higo

type Handlers []Handler
type Handler func(*Context)
