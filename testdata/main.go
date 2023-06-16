//go:build err4

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

func fn() (qTry error) {
	qTry = errors.New("e")
	return
}
