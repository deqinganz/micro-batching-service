package http

import (
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"micro-batching/api"
	"net/http"
)

type Handlers struct {
}

func (s *Handlers) GetBatchFrequency(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Handlers) PostBatchFrequency(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Handlers) GetBatchSize(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Handlers) UpdateBatchSize(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Handlers) PostJob(c *gin.Context) {
	var job api.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (*Handlers) GetJobId(c *gin.Context, id openapitypes.UUID) {
}
