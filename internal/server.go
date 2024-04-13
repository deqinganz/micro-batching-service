package internal

import (
	"github.com/deqinganz/batching-api/api"
	"github.com/deqinganz/micro-batching"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"log"
)

func SetupHandler(batching *batching.Batching) {
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
