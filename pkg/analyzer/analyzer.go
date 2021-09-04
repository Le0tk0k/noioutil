package analyzer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// Analyzer is the noioutil analyzer.
var Analyzer = &analysis.Analyzer{
	Name:     "noioutil",
	Doc:      "noioutil finds io/ioutil package",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}
