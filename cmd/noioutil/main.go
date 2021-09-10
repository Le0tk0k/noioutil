package main

import (
	"github.com/le0tk0k/noioutil"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noioutil.Analyzer) }
