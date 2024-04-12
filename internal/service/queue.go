package service

import (
	"container/ring"
	. "micro-batching/api"
)

type Queue interface {
	ring.Ring
	Enqueue(job Job) error
	Dequeue() (Job, error)
}
