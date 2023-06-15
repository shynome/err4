//go:build !err4

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
