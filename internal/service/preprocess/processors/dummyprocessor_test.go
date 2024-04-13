package processors

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"testing"
)

func TestDummyProcessor(t *testing.T) {
	processor := &DummyProcessor{}

	jobs := processor.Process([]Job{{}, {}})

	assert.Equal(t, []Job{{}, {}}, jobs)
}
