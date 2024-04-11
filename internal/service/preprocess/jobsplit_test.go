package preprocess

import (
	"github.com/stretchr/testify/assert"
	. "micro-batching/api"
	"testing"
)

func TestSplit(t *testing.T) {
	jobMap := Split([]Job{{Type: "A"}, {Type: "B"}, {Type: "A"}})

	assert.Equal(t, 2, len(jobMap["A"]))
	assert.Equal(t, 1, len(jobMap["B"]))
}
