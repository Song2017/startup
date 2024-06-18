package service

import (
	"github.com/gin-gonic/gin"
)

type GinProcess interface {
	Input()
	Process()
	Output()

	Invoke(*gin.Context, ...OperationBase) bool
}
type Operation func(*gin.Context)

type ServiceGin struct {
	ServiceName string
	ServiceBase
}

func (p *ServiceGin) Input(c *gin.Context) {
	p.ServiceBase.Input()
}

func (p *ServiceGin) Output(c *gin.Context) {
	if p.Response[p.KeyError] != nil {
		c.AbortWithStatusJSON(400, p.Response[p.KeyError])
		return
	}
}

func (p *ServiceGin) Invoke(c *gin.Context, funcs ...Operation) bool {
	// Invoke functions to process Request
	for _, f := range funcs {
		f(c)
		// return diretly if error
		if p.Response[p.KeyError] != nil {
			p.Output(c)
			return false
		}
	}
	return true
}
