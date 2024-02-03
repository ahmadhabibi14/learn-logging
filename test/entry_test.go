package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello logging")
	logger.WithField("username", "haaabibich").Info("Hello logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "ahmadhabibi14")
	entry.Info("Hello entry")
}
