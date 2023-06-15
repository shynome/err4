package err4_test

import (
	"errors"
	"testing"

	"github.com/lainio/err2/assert"
	"github.com/shynome/err4"
)

func TestCheck(t *testing.T) {
	var err error

	var (
		e1 = errors.New("1")
		e2 = errors.New("2")
		e3 = errors.New("3")
		e4 = errors.New("4")
		e5 = errors.New("5")
		e6 = errors.New("6")
		e7 = errors.New("7")
		e8 = errors.New("8")
		e9 = errors.New("9")
	)
	var a1, a2, a3, a4, a5, a6, a7, a8, a9 any

	_ = err4.Check(e1)(&err)
	assert.Equal(err, e1)
	err = nil

	a1, _ = err4.Check1(1, e1)(nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(err, e1)
	err = nil

	a1, a2, _ = err4.Check2(1, 2, e2)(nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(err, e2)
	err = nil

	a1, a2, a3, _ = err4.Check3(1, 2, 3, e3)(nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(err, e3)
	err = nil

	a1, a2, a3, a4, _ = err4.Check4(1, 2, 3, 4, e4)(nil, nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(a4, 4)
	assert.Equal(err, e4)
	err = nil

	a1, a2, a3, a4, a5, _ = err4.Check5(1, 2, 3, 4, 5, e5)(nil, nil, nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(a4, 4)
	assert.Equal(a5, 5)
	assert.Equal(err, e5)
	err = nil

	a1, a2, a3, a4, a5, a6, _ = err4.Check6(1, 2, 3, 4, 5, 6, e6)(nil, nil, nil, nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(a4, 4)
	assert.Equal(a5, 5)
	assert.Equal(a6, 6)
	assert.Equal(err, e6)
	err = nil

	a1, a2, a3, a4, a5, a6, a7, _ = err4.Check7(1, 2, 3, 4, 5, 6, 7, e7)(nil, nil, nil, nil, nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(a4, 4)
	assert.Equal(a5, 5)
	assert.Equal(a6, 6)
	assert.Equal(a7, 7)
	assert.Equal(err, e7)
	err = nil

	a1, a2, a3, a4, a5, a6, a7, a8, _ = err4.Check8(1, 2, 3, 4, 5, 6, 7, 8, e8)(nil, nil, nil, nil, nil, nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(a4, 4)
	assert.Equal(a5, 5)
	assert.Equal(a6, 6)
	assert.Equal(a7, 7)
	assert.Equal(a8, 8)
	assert.Equal(err, e8)
	err = nil

	a1, a2, a3, a4, a5, a6, a7, a8, a9, _ = err4.Check9(1, 2, 3, 4, 5, 6, 7, 8, 9, e9)(nil, nil, nil, nil, nil, nil, nil, nil, nil, &err)
	assert.Equal(a1, 1)
	assert.Equal(a2, 2)
	assert.Equal(a3, 3)
	assert.Equal(a4, 4)
	assert.Equal(a5, 5)
	assert.Equal(a6, 6)
	assert.Equal(a7, 7)
	assert.Equal(a8, 8)
	assert.Equal(a9, 9)
	assert.Equal(err, e9)
	err = nil

	a1, a2, a3, a4, a5, a6, a7, a8, a9, _ = err4.Check9(e1, e2, e3, e4, e5, e6, e7, e8, e9, e9)(&err, nil, nil, nil, nil, nil, nil, nil, nil, &err)
	assert.Equal(a1.(error), e1)
	assert.Equal(a2.(error), e2)
	assert.Equal(a3.(error), e3)
	assert.Equal(a4.(error), e4)
	assert.Equal(a5.(error), e5)
	assert.Equal(a6.(error), e6)
	assert.Equal(a7.(error), e7)
	assert.Equal(a8.(error), e8)
	assert.Equal(a9.(error), e9)
	assert.Equal(err, e1)
	err = nil

}
