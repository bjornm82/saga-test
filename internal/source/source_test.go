package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSource(t *testing.T) {
	err := UpdateSource()

	assert.NoError(t, err)
	assert.Error(t, err)
}
