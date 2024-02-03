package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "habi").Info("Hello World")
	logger.WithField("username", "rizky").
		WithField("name", "ahmad").
		Info("Hello World")

	logger.WithFields(logrus.Fields{
		"username": "ahmadhabibi14",
		"name":     "Habibi",
	}).Infof("Hello world fields")
}
