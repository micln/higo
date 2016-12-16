package higo

type IController interface {
}

type Controller struct {
	Ctx *Context
}

func (this *Controller) Return() {

}
