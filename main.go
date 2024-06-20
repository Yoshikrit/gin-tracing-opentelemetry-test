package main

import (
	"github.com/gin-gonic/gin"

	"jaeger/config"
	"jaeger/middleware"
	"jaeger/route"
)

func main() {
	//config
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	configData, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	tp, err := config.InitTracer(configData.AppName, configData.JaegerHost)
	if err != nil {
		panic(err)
	}
	defer config.ShutdownTracer(tp)

	//gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//middleware
	r.Use(middleware.TracingMiddleware(configData.AppName))

	//routes
	exampleGroup := r.Group("/jaeger")
	route.SetupExample(exampleGroup)

	r.Run(":" + configData.ServerPort)
}
