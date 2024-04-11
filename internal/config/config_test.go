package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfig(t *testing.T) {
	expectedConfig := RunConfig{
		BatchSize: 10,
		Frequency: 1,
	}

	readConfig, err := ReadConfig()
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, readConfig)
}
