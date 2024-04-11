package service

import (
	"github.com/google/uuid"
	. "micro-batching/api"
	"micro-batching/internal/config"
)

type Batching struct {
	config config.RunConfig
	queue  []Job
}

func (b *Batching) Take(jobRequest JobRequest) Job {
	var job Job
	job.Id = uuid.New()
	job.Status = QUEUED
	job.Type = jobRequest.Type
	job.Name = jobRequest.Name
	job.Params = Job_Params(jobRequest.Params)

	b.queue = append(b.queue, job)

	return job
}
