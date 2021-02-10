package gRouter

import "reflect"

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

func (handlers HandlersChain) Names() []string{
	list := make([]string, len(handlers))
	for index, handler := range handlers {
		list[index] = reflect.TypeOf(handler).String()
	}
	return list
}