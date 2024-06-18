package main

import (
	"log"
	"net/http"
	_init "startup/_core/_init"
	_middleware "startup/_middleware"

	"github.com/gin-gonic/gin"
)

// go run tests/test_main.go
// http://localhost:9000/swagger/index.html
func addRoutes(router *gin.Engine) {
	// endpoints
	order_group := router.Group("/order")
	{
		order_group.Use(
			_middleware.LoggerMiddleware(_init.O_SERVER_CONFIG.ServerName),
			_middleware.AuthorizationMiddleware(
				_init.O_SERVER_CONFIG.SecurityKey, _init.O_SERVER_CONFIG.SecurityValue),
		)
		order_group.POST(
			"/v1/internal/ranger/create",
			func(c *gin.Context) {
				c.String(http.StatusOK, "root page")
			},
		)
	}

	odoo_group := router.Group("/order")
	{
		odoo_group.Use(
			_middleware.LoggerMiddleware(_init.O_SERVER_CONFIG.ServerName),
			_middleware.AuthorizationMiddleware(
				_init.O_SERVER_CONFIG.SecurityKey, _init.O_SERVER_CONFIG.SecurityValue),
		)
	}

}

func main() {
	// init_project.InitProjectConfig()
	_init.InitServerConfig()
	router := _init.InitRouter(nil)
	_init.InitSwagger(router, "./resources")

	addRoutes(router)

	// init server
	s := _init.InitServer("0.0.0.0:9000", router)
	log.Fatal(s.ListenAndServe().Error())
}
