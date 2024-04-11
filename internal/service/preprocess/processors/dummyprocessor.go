package processors

import (
	. "micro-batching/api"
)

type DummyProcessor struct {
}

func (m *DummyProcessor) Process(jobs []Job) []Job {
	return jobs
}
