package controller

import (
	"github.com/micln/higo/context"
	"github.com/micln/higo/exception"
)

type IController interface {
}

type Controller struct {
	Ctx         *context.Context
}

func (c *Controller) Return() {

}

func (c *Controller) ThrowException(exception exception.IException) {

}
