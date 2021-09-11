package noioutil

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "noioutil finds files using the io/ioutil package."

// Analyzer is the noioutil analyzer.
var Analyzer = &analysis.Analyzer{
	Name: "noioutil",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			fileName := n.Name.Name
			for _, f := range n.Decls {
				switch decl := f.(type) {
				case *ast.GenDecl:
					if decl.Tok == token.IMPORT {
						checkImport(pass, decl, fileName)
					}
				}
			}
		}
	})

	return nil, nil
}

func checkImport(pass *analysis.Pass, decl *ast.GenDecl, fileName string) {
	for _, spec := range decl.Specs {
		if spec, ok := spec.(*ast.ImportSpec); ok {
			if spec.Path.Value == `"io/ioutil"` {
				pass.Reportf(spec.Path.Pos(), "\"io/ioutil\" package is used")
			}
		}
	}
}
