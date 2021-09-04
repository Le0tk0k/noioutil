package main

import (
	"github.com/Le0tk0k/noioutil"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noioutil.Analyzer) }
