package service

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"testing"
)

func TestProcess(t *testing.T) {
	jobs := []Job{{Status: QUEUED}, {Status: QUEUED}}
	Process(jobs)

	assert.Equal(t, []Job{{Status: SUBMITTED}, {Status: SUBMITTED}}, jobs)
}
