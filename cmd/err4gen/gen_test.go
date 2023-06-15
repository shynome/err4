package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
)

func TestGen(t *testing.T) {
	var output bytes.Buffer
	try.To(gen("../../testdata/main.go", nil, &output))
	f := try.To1(os.Open("../../testdata/main_err4.go"))
	b := try.To1(io.ReadAll(f))
	assert.Equal(output.String(), string(b))
}
