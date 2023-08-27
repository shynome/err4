//go:build ierr

package main

import (
	"errors"
)

func main() {
	var ierr error
	var ierr2 error
	_, _ = ierr, ierr2

	ierr = errors.New("e")

	var e1, e2 error
	_, ierr, ierr2 = 7, e1, e2
}

func fn() (ierr error) {
	ierr = errors.New("e")
	return
}

func fn2() (ierr error) {
	ierr = errors.New("e")
	return
}

func fn3() (ierr error) {
	if ierr = errors.New(""); ierr != nil {
		return
	}
	return
}
