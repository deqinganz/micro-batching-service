package http

import (
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"micro-batching/api"
	"micro-batching/internal/service"
	"net/http"
)

type Handlers struct {
	batching *service.Batching
}

func (h *Handlers) GetBatchFrequency(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"frequency": h.batching.GetFrequency().Frequency})
}

func (h *Handlers) PostBatchFrequency(c *gin.Context) {
	var frequency api.BatchFrequency
	if err := c.BindJSON(&frequency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		h.batching.SetFrequency(frequency)
		c.JSON(http.StatusOK, gin.H{"frequency": frequency.Frequency})
	}
	h.batching.Restart()
}

func (h *Handlers) GetBatchSize(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"batch size": h.batching.GetBatchSize().BatchSize})
}

func (h *Handlers) UpdateBatchSize(c *gin.Context) {
	var batchSize api.BatchSize
	if err := c.BindJSON(&batchSize); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		h.batching.SetBatchSize(batchSize)
		c.JSON(http.StatusOK, gin.H{"batchSize": batchSize.BatchSize})
	}
	h.batching.Restart()
}

func (h *Handlers) PostJob(c *gin.Context) {
	var jobRequest api.JobRequest
	if err := c.BindJSON(&jobRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		job := h.batching.Take(jobRequest)
		c.JSON(http.StatusOK, gin.H{"job": job})
	}
}

func (h *Handlers) GetJobById(c *gin.Context, id openapitypes.UUID) {
	job, err := h.batching.JobInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"job": job})
	}
}
