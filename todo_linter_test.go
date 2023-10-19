package todo_linter_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/mrymam/todo_linter"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, todo_linter.Analyzer, "a")
}
