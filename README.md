### A implemention for golang Error-Handling

for this proposal [proposal: Go 2: Error-Handling Paradigm with ?err Grammar Sugar](https://github.com/golang/go/issues/60779)

use `var qTerr error` instedof `var err error`, `qTerr` = `?err` -> `qT` = `?`

### how to use

source code

```golang
package main

import "errors"

func main() {
	var qTerr error
	var qTerr2 error
	_, _ = qTerr, qTerr2

	qTerr = errors.New("e")

	var e1, e2 error
	_, qTerr, qTerr2 = 7, e1, e2
}

```

after transpile with `go run ./cmd/err4gen/ -f ./testdata/main.go`

```golang
// Code generated github.com/shynome/err4 DO NOT EDIT
package main

import (
	"errors"
	"github.com/shynome/err4"
)

func main() {
	var qTerr error
	var qTerr2 error
	_, _ = qTerr, qTerr2

	_ = err4.Check(errors.New("e"))(&qTerr)
	if qTerr != nil {
		return
	}

	var e1, e2 error
	_, _, _ = err4.Check2(7, e1, e2)(nil, &qTerr, &qTerr2)
	if qTerr != nil {
		return
	}
	if qTerr2 != nil {
		return
	}
}

```
