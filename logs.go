package higo

import "log"

type ILogger interface {
Info(interface{})
}

type Logger struct {

}

func (lg *Logger) Info(v interface{})  {
	log.Println(v)
}