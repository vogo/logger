package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	assert.Equal(t, "helloworld", fmt.Sprint("hello", "world"))
	logger := DefaultLogger

	logger.Info("hello", "world")
	logger.Warn("hello", "world", "test")
	logger.Error("hello", "world", "test")
	logger.Debug("hello", "world", "test")

	logger.Infof("hello %s", "world")
	logger.Warnf("hello %s %s", "world", "test")
	logger.Errorf("hello %s %s", "world", "test")
	logger.Debugf("hello %s %s", "world", "test")
}
