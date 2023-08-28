package main

import (
	"fmt"
	"testing"
)

func TestErr4Path(t *testing.T) {
	testcases := [][]string{
		{"dir/a/a.go", "dir/a/a_ierr.go"},
		{"dir/a/a_js.go", "dir/a/a_ierr_js.go"},
		{"dir/a/a_test.go", "dir/a/a_ierr_test.go"},
		{"../../dir/a/a_test.go", "../../dir/a/a_ierr_test.go"},
	}
	for _, v := range testcases {
		a, b := v[0], v[1]
		aa := err4Path(a)
		if aa != b {
			err := fmt.Errorf("want %s, got %s", b, aa)
			t.Error(err)
			return
		}
	}
}
