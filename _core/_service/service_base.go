package service

import (
	"context"
)

type BaseProcess interface {
	Input()
	Process()

	Invoke(*context.Context, ...OperationBase) bool
}

type OperationBase func(*context.Context)

type ServiceBase struct {
	Request map[string]interface{}
	Body    map[string]interface{}

	Response  map[string]interface{} // code, result, error
	KeyResult string
	KeyError  string
}

func (p *ServiceBase) Input() {
	// return err or failed check result
	p.KeyResult = "result"
	p.KeyError = "error"

	p.Request = make(map[string]interface{})
	p.Response = map[string]interface{}{}

}

func (p *ServiceBase) Process() {
	// business process

}

func (p *ServiceBase) Invoke(c *context.Context, funcs ...OperationBase) bool {
	for _, f := range funcs {
		f(c)
		if p.Response[p.KeyError] != nil {
			return false
		}
	}
	return true
}
