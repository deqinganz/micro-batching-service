package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"testing"
)

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(Job{Name: "j1"})
	q.Enqueue(Job{Name: "j2"})
	q.Enqueue(Job{Name: "j3"})
	assert.Equal(t, 3, q.Size())

	jobs := q.Dequeue(2)
	assert.NotNil(t, jobs)
	assert.Equal(t, "j1", jobs[0].Name)
	assert.Equal(t, "j2", jobs[1].Name)
	assert.Equal(t, 1, q.Size())

	jobs = q.Dequeue(2)
	assert.NotNil(t, jobs)
	assert.Equal(t, 1, len(jobs))
	assert.Equal(t, "j3", jobs[0].Name)

	q.Enqueue(Job{Name: "j4"})
	jobs = q.Dequeue(10)
	assert.NotNil(t, jobs)
	assert.Equal(t, 1, len(jobs))
	assert.Equal(t, "j4", jobs[0].Name)

	jobs = q.Dequeue(10)
	assert.Equal(t, 0, len(jobs))
}

func TestFind(t *testing.T) {
	q := &Queue{}
	id := uuid.New()
	q.Enqueue(Job{Id: uuid.New(), Name: "Job1"})
	q.Enqueue(Job{Id: id, Name: "Job2"})
	q.Enqueue(Job{Id: uuid.New(), Name: "Job3"})

	job, err := q.Find(id)
	assert.NoError(t, err)
	assert.NotNil(t, job)
	assert.Equal(t, "Job2", job.Name)

	_, err = q.Find(uuid.New())
	assert.Error(t, fmt.Errorf("job not found"), err)
}
