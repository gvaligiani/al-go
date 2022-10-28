package test

import (
	"fmt"
	"testing"
	"time"
	"reflect"

	"go.uber.org/zap"
)

// RunTestCases runs a set of test cases
func RunTestCases[TestCase any](t *testing.T,testCases map[string]TestCase,runTestCase func(t *testing.T, logger *zap.Logger, testCase TestCase)) {
	t.Helper()
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Helper()
			
			logger := NewLogger()
			logger.Info(fmt.Sprintf(" ----- %s/%v ",reflect.TypeOf(testCase).PkgPath(),t.Name()))
			start := time.Now()
			defer logger.Info(fmt.Sprintf(" >>> test completed in %v ",time.Since(start)))

			runTestCase(t,NewTestCaseLogger(testName),testCase)
		})
	}
}
