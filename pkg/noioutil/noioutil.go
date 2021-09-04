package noioutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	importStartFlag = []byte(`
import (
`)
	importEndFlag = []byte(`
)
`)
)

func Run(fileName string) error {
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
	end := bytes.Index(src[start:], importEndFlag) + start
	pkgs := src[start+len(importStartFlag) : end]

	var useIoutil bool
	if strings.Contains(string(pkgs), "io/ioutil") {
		useIoutil = true
	}

	if useIoutil {
		return fmt.Errorf("use \"io/ioutil\" package in %s", fileName)
	}
	return nil
}
