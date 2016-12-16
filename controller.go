package higo

import "github.com/micln/higo/exception"

type IController interface {
}

type Controller struct {
	Ctx *Context
}

func (c *Controller) Return() {

}

func (c *Controller) ThrowException(exception exception.IException) {

}