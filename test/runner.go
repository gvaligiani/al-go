package test

import (
	"testing"
	"time"
	
	"go.uber.org/zap"
)

// RunTestCases runs a set of test cases
func RunTestCases[TestCase any](t *testing.T,testCases map[string]TestCase,runTestCase func(testName string, logger *zap.Logger, testCase TestCase)) {
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			logger := NewTestCaseLogger(testName)
			logger.Info(" ---------- start ---------- ")
			start := time.Now()
			defer logger.Info(" ---------- end ---------- ", zap.Duration("time", time.Since(start)))

			runTestCase(testName,logger,testCase)
		})
	}	
}

