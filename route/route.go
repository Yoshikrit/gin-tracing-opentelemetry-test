package route

import (
	"jaeger/handler"
	"jaeger/repository"
	"jaeger/service"

	"github.com/gin-gonic/gin"
)

func SetupExample(g *gin.RouterGroup) {
	exampleRepository := repository.NewExampleRepository()
	exampleService := service.NewExampleService(exampleRepository)
	exampleHandler := handler.NewExampleHandler(exampleService)

	g.GET("/:id", exampleHandler.Post)
	g.GET("/", exampleHandler.Get)
}
