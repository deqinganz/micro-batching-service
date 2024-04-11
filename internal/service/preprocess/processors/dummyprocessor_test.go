package processors

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"micro-batching/internal/service/preprocess"
	"testing"
)

func TestDummyProcessor(t *testing.T) {
	var processor preprocess.Processor
	processor = &DummyProcessor{}

	jobs := processor.Process([]Job{{}, {}})

	assert.Equal(t, []Job{{}, {}}, jobs)
}
