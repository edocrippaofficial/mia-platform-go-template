package router

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	log, _ := test.NewNullLogger()
	router := NewRouter(log)
	assert.NotNil(t, router, "Router should not be nil")
}
