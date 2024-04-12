package service

import (
	"github.com/google/uuid"
	"log"
	. "micro-batching/api"
	"micro-batching/internal/config"
)

type Batching struct {
	config         config.RunConfig
	queue          Queue
	batchProcessor BatchProcessor
}

func NewBatching() (Batching, error) {
	c, err := config.ReadConfig()
	if err != nil {
		return Batching{}, err
	}

	return Batching{
		config: c,
		queue:  Queue{},
	}, nil

}

func NewBatchingWithConfig(c config.RunConfig) Batching {
	return Batching{
		config: c,
		queue:  Queue{},
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

	b.queue.Enqueue(job)

	return job
}

// JobInfo returns the job by the given id
func (b *Batching) JobInfo(id uuid.UUID) (Job, error) {
	j, err := b.queue.Find(id)
	if err != nil {
		return Job{}, err
	}

	return j, nil
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

// TODO 1. 收到新的frequency之后要把这个post中断，重新来\
func (b *Batching) Post() {
	if b.queue.Size() == 0 {
		log.Print("No jobs to process")
		return
	}

	jobs := b.queue.Dequeue(b.config.BatchSize)

	b.batchProcessor.Process(jobs)
}
