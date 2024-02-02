package test

import (
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	multiOut := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(multiOut)

	logger.Info("Hello logging")
	logger.Warn("Hello logging")
	logger.Error("Hello logging")
}
