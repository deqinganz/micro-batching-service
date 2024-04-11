package service

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"testing"
)

func TestProcess(t *testing.T) {
	jobs := []Job{{Status: NEW}, {Status: NEW}}
	Process(jobs)

	assert.Equal(t, []Job{{Status: SUBMITTED}, {Status: SUBMITTED}}, jobs)
}
