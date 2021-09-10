package a

import (
	"io/ioutil" // want "\"io/ioutil\" package is used"
)

func f() {
	_ = ioutil.Discard
}
