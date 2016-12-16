package exception

import "fmt"

//	@implements IException
type DefaultException struct {
	Code    int
	Message string
}

func (e *DefaultException) Response() string {
	return fmt.Sprintf("%d : %s", e.Code, e.Message)
}
