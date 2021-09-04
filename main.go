package main

import (
	"flag"
	"go/scanner"
	"os"
	"path/filepath"
	"strings"

	"github.com/Le0tk0k/noioutil/pkg/noioutil"
)

var exitCode = 0

func parseFlags() []string {
	flag.Parse()
	return flag.Args()
}

func report(err error) {
	if err == nil {
		return
	}
	scanner.PrintError(os.Stderr, err)
	exitCode = 1
}

func isGoFile(f os.FileInfo) bool {
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isGoFile(f) {
		err = noioutil.Run(path)
	}
	if err != nil {
		report(err)
	}
	return nil
}

func walkDir(path string) {
	filepath.Walk(path, visitFile)
}

func main() {
	paths := parseFlags()

	for _, path := range paths {
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsDir():
			walkDir(path)
		default:
			if err := noioutil.Run(path); err != nil {
				report(err)
			}
		}
	}
	os.Exit(exitCode)
}
