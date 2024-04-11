package http

import (
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"micro-batching/api"
	"micro-batching/internal/service"
	"net/http"
)

type Handlers struct {
	batching service.Batching
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
	var jobRequest api.JobRequest
	if err := c.BindJSON(&jobRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"job": s.batching.Take(jobRequest)})
	}
}

func (*Handlers) GetJobId(c *gin.Context, id openapitypes.UUID) {
}
