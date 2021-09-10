# noioutil

`noioutil` finds files using the "io/ioutil" package.

## Installation

```
go install github.com/le0tk0k/noioutil/cmd/noioutil@latest
```

## Usage

```
go vet -vettool=$(which noioutil) ./...
```

## Example

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("foo.go")
	fmt.Println(string(bytes))
}
```

```
$ go vet -vettool=$(which noioutil) ./...
noioutil: "io/ioutil" package is used in ./main.go
```

