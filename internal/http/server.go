package http

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"log"
	"micro-batching/api"
	"micro-batching/internal/service"
)

func SetupHandler(batching *service.Batching) {
	r := gin.Default()

	swagger, _ := api.GetSwagger()
	r.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(r, &Handlers{batching})

	batching.Start()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server")
		return
	}
}
