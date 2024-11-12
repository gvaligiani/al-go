package test

import (
	"go.uber.org/zap"
)

// NewZapLogger builds test zap logger
func NewZapLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Development = false
	logger, _ := config.Build()
	return logger
}

// NewTestCaseLogger builds test logger for a test case
func NewTestCaseLogger(testName string) *zap.Logger {
	return NewZapLogger().With(zap.String("test", testName))
}
