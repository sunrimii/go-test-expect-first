package analyzer

import (
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "expectfirst",
	Doc:      "Checks the expect-liked argument should be first replaced and then actual-liked argument in test assertion",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// check file using testify/require only
	if !analysisutil.Imported("github.com/stretchr/testify/require", pass) {
		return nil, nil
	}

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		callExpr := node.(*ast.CallExpr)
		selectorExpr := callExpr.Fun.(*ast.SelectorExpr)
		if selectorExpr.X.(*ast.Ident).Name != "require" {
			return
		}
		if selectorExpr.Sel.Name != "Equal" {
			return
		}

		args := callExpr.Args
		if len(args) != 3 {
			return
		}
		if args[0].(*ast.Ident).Name != "t" {
			return
		}

		secondArg, ok := args[1].(*ast.Ident)
		if ok && isActualLike(secondArg.Name) {
			pass.Reportf(node.Pos(), "the actual-like variable name must comes after the expect-like variable name")
			return
		}

		thirdArg, ok := args[2].(*ast.Ident)
		if ok && isExpectLike(thirdArg) {
			pass.Reportf(node.Pos(), "the expected-like variable name must comes before the actual-like variable name")
			return
		}
	})

	return nil, nil
}

func isActualLike(name string) bool {
	name = strings.ToLower(name)
	for _, actaulLike := range []string{"actual", "got"} {
		if strings.HasPrefix(name, actaulLike) || strings.HasSuffix(name, actaulLike) {
			return true
		}
	}
	return false
}

func isExpectLike(arg *ast.Ident) bool {
	name := strings.ToLower(arg.Name)
	for _, expectLike := range []string{"expect", "want"} {
		if strings.HasPrefix(name, expectLike) || strings.HasSuffix(name, expectLike) {
			return true
		}
	}
	return false
}
