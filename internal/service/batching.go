package service

import (
	"errors"
	"github.com/google/uuid"
	. "micro-batching/api"
	"micro-batching/internal/config"
)

type Batching struct {
	config config.RunConfig
	queue  []Job
}

func NewBatching() (Batching, error) {
	c, err := config.ReadConfig()
	if err != nil {
		return Batching{}, err
	}

	return Batching{
		config: c,
	}, nil

}

func NewBatchingWithConfig(config config.RunConfig) Batching {
	return Batching{
		config: config,
	}
}

// Take creates a new job and adds it to the queue
func (b *Batching) Take(jobRequest JobRequest) Job {
	job := Job{
		Id:     uuid.New(),
		Status: QUEUED,
		Type:   jobRequest.Type,
		Name:   jobRequest.Name,
		Params: Job_Params(jobRequest.Params),
	}

	b.queue = append(b.queue, job)

	return job
}

func (b *Batching) JobInfo(id uuid.UUID) (Job, error) {
	for _, job := range b.queue {
		if job.Id == id {
			return job, nil
		}
	}

	return Job{}, errors.New("job not found")
}

func (b *Batching) GetFrequency() BatchFrequency {
	return BatchFrequency{
		Frequency: b.config.Frequency,
	}
}

func (b *Batching) GetBatchSize() BatchSize {
	return BatchSize{
		BatchSize: b.config.BatchSize,
	}
}

func (b *Batching) SetFrequency(frequency BatchFrequency) {
	b.config.Frequency = frequency.Frequency
}

func (b *Batching) SetBatchSize(batchSize BatchSize) {
	b.config.BatchSize = batchSize.BatchSize
}
