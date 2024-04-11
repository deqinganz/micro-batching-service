package http

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"log"
	"micro-batching/api"
	"micro-batching/internal/service"
	"os"
)

func SetupHandler() {
	r := gin.Default()

	swagger, _ := api.GetSwagger()
	r.Use(middleware.OapiRequestValidator(swagger))

	batching, err := service.NewBatching()
	if err != nil {
		path, _ := os.Getwd()
		log.Fatalf("creating batching service failed: %v in %s", err, path)
		return
	}

	api.RegisterHandlers(r, &Handlers{batching})
	r.Run(":8080")
}
