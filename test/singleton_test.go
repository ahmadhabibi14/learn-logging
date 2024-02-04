package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello world")
	logrus.Warn("Hello world")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello world")
	logrus.Warn("Hello world")
}
