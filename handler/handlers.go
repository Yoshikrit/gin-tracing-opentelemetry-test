package handler

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"

	"jaeger/service"

	"net/http"
)

var tracer = otel.Tracer("jaeger-gin-example")

type exampleHandler struct {
	exampleSrv service.ExampleService
}

func NewExampleHandler(exampleSrv service.ExampleService) *exampleHandler {
	return &exampleHandler{exampleSrv: exampleSrv}
}

func (h *exampleHandler) Post(ctx *gin.Context) {
	_, span := tracer.Start(ctx.Request.Context(), "Post_Handler")
	defer span.End()

	id := ctx.Param("id")
	response := h.exampleSrv.Post(&id)

	ctx.JSON(http.StatusOK, response)
}

func (h *exampleHandler) Get(ctx *gin.Context) {
	_, span := tracer.Start(ctx.Request.Context(), "Get_Handler")
	defer span.End()

	response := h.exampleSrv.Get()

	ctx.JSON(http.StatusOK, response)
}
