package service

import (
	"container/ring"
	"errors"
	"github.com/google/uuid"
	"log"
	. "micro-batching/api"
	"micro-batching/internal/config"
)

type Batching struct {
	config         config.RunConfig
	queue          *ring.Ring
	queueSize      int // ring.Len() is O(n) so we keep track of the size to have better time complexity
	batchProcessor BatchProcessor
}

func NewBatching() (Batching, error) {
	c, err := config.ReadConfig()
	if err != nil {
		return Batching{}, err
	}

	return Batching{
		config:    c,
		queue:     ring.New(c.QueueSize),
		queueSize: 0,
	}, nil

}

func NewBatchingWithConfig(c config.RunConfig) Batching {
	return Batching{
		config:    c,
		queue:     ring.New(c.QueueSize),
		queueSize: 0, //TODO remove?
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

	b.queue.Value = job
	b.queue = b.queue.Next()
	b.queueSize++

	return job
}

// JobInfo returns the job by the given id
func (b *Batching) JobInfo(id uuid.UUID) (Job, error) {
	j := Job{}
	b.queue.Do(func(job interface{}) {
		if job != nil && job.(Job).Id == id {
			j = job.(Job)
			return
		}
	})
	if j.Id != uuid.Nil {
		return j, nil
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

// TODO 1. 收到新的frequency之后要把这个post中断，重新来
// 2. 取出指定大小的batchsize去调用，然后把queue里面前面的部分去掉
func (b *Batching) Post() {
	if b.queueSize == 0 {
		log.Print("No jobs to process")
		return
	}

	//b.batchProcessor.Process(b.queue)
}
