package err4_test

import (
	"errors"
	"testing"

	"github.com/lainio/err2/assert"
	"github.com/shynome/err4"
)

func TestHandle(t *testing.T) {
	var err error
	a := 1
	err4.Handle(&err)(func() {
		a = 2
	})
	assert.Equal(a, 1)

	var err2 error = errors.New("b")
	b := 1
	err4.Handle(&err2)(func() {
		b = 2
	})
	assert.Equal(b, 2)

}
