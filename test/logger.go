package test

import (
	"go.uber.org/zap"
)

// NewLogger builds test logger
func NewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Development = false
	logger, _ := config.Build()
	return logger
}

// NewTestCaseLogger builds test logger for a test case
func NewTestCaseLogger(testName string) *zap.Logger {
	return NewLogger().With(zap.String("test", testName))
}
