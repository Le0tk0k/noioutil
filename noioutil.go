package noioutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
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
	var fileNames []string
	for _, f := range pass.Files {
		pos := pass.Fset.PositionFor(f.Pos(), false)
		fileNames = append(fileNames, pos.Filename)
	}

	for _, f := range fileNames {
		err := processFile(f)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

var (
	importStartFlag = []byte(`
import (
`)
	importEndFlag = []byte(`
)
`)
)

func processFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	start := bytes.Index(src, importStartFlag)
	if start < 0 {
		return nil
	}
	end := bytes.Index(src[start:], importEndFlag) + start
	pkgs := src[start+len(importStartFlag) : end]

	var useIoutil bool
	if strings.Contains(string(pkgs), "io/ioutil") {
		useIoutil = true
	}

	if useIoutil {
		return fmt.Errorf("\"io/ioutil\" package is used in %s", fileName)
	}
	return nil
}
