package transpile

import (
	"io"
	"os"
	"testing"

	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
)

func TestGen(t *testing.T) {
	output, _ := try.To2(Transform("../../testdata/main.go", nil))
	f := try.To1(os.Open("../../testdata/main_ierr.go"))
	b := try.To1(io.ReadAll(f))
	assert.Equal(output.String(), string(b))
}
