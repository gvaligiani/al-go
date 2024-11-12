package test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"go.uber.org/zap"
)

type testKey struct{}

var TestKey = testKey{}

type TestCaseFn[TestCase any] func(t *testing.T, cxt context.Context, logger *zap.Logger, testCase TestCase)

// RunTestCases runs a set of test cases
func RunTestCases[TestCase any](t *testing.T, testCases map[string]TestCase, testCaseFn TestCaseFn[TestCase]) {
	t.Helper()

	ctx := context.Background()
	logger := NewZapLogger()

	testPattern := os.Getenv("TEST_PATTERN")

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Helper()

			pkgPath := reflect.TypeOf(testCase).PkgPath()
			splits := strings.Split(pkgPath, "/")
			last := splits[len(splits)-1]

			if testPattern != "" {
				if match, _ := regexp.MatchString(testPattern, testName); !match {
					logger.Info(fmt.Sprintf(" --- OFF --- %s :: %v", last, t.Name()))
					return
				}
			}

			logger.Info(fmt.Sprintf(" ----------- %s :: %v ", last, t.Name()))
			start := time.Now()
			defer func() {
				logger.Info(fmt.Sprintf(" >>> test completed in %v ", time.Since(start)))
			}()

			testCaseFn(
				t,
				context.WithValue(ctx, TestKey, testName),
				logger.With(zap.String("test", testName)),
				testCase,
			)
		})
	}
}
