package test

import (
	"github.com/stretchr/testify/assert"
	"go-api/simple"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
