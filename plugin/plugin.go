// This must be package main
package main

import (
	"expect-first/pkg/analyzer"

	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.Analyzer}
}

var AnalyzerPlugin analyzerPlugin
