package processors

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"micro-batching/internal/service"
	"testing"
)

func TestDummyProcessor(t *testing.T) {
	var processor service.Processor
	processor = &DummyProcessor{}

	jobs := processor.Process([]Job{{}, {}})

	assert.Equal(t, []Job{{}, {}}, jobs)
}
