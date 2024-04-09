package http

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"micro-batching/api"
)

func SetupHandler() {
	r := gin.Default()

	swagger, _ := api.GetSwagger()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(r, &Handlers{})

	r.Run(":8080")
}
