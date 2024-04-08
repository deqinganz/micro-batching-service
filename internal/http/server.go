package http

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"micro-batching/api"
	"net/http"
)

type Server struct {
}

func (s *Server) GetBatchFrequency(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostBatchFrequency(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetBatchSize(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateBatchSize(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostJob(c *gin.Context) {
	var job api.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (*Server) GetJobId(c *gin.Context, id openapi_types.UUID) {
}

func SetupHandler() {
	r := gin.Default()

	swagger, _ := api.GetSwagger()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(r, &Server{})

	r.Run(":8080")
}
