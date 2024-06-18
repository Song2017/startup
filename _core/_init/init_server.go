package initial

import (
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server interface {
	ListenAndServe() error
}

func InitSwagger(router *gin.Engine, staticPath string) *gin.Engine {
	// swagger UI
	// http://localhost:9000/swagger/index.html
	if staticPath == "" {
		staticPath = "./_core/resources"
	}
	router.Static("/static", staticPath)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler, ginSwagger.URL("/static/openapi.yaml")))
	return router
}

func InitRouter(router *gin.Engine) *gin.Engine {
	if router == nil {
		router = gin.New()
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "root page")
	})
	return router
}

func InitServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
