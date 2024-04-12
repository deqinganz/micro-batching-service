package http

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"micro-batching/api"
	"micro-batching/internal/service"
)

func SetupHandler(batching *service.Batching) {
	r := gin.Default()

	swagger, _ := api.GetSwagger()
	r.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(r, &Handlers{batching})
	r.Run(":8080")
}
