package err4

func CommonCheck(errs []*error) func(vars ...any) {
	return func(vars ...any) {
		for i, err := range errs {
			if err == nil || *err != nil {
				continue
			}
			if e := vars[i]; e != nil {
				*err = e.(error)
			}
		}
	}
}

func Check[A1 any](a1 A1) func(errs ...*error) A1 {
	return func(errs ...*error) A1 {
		CommonCheck(errs)(a1)
		return a1
	}
}

func Check1[A1, A2 any](a1 A1, a2 A2) func(errs ...*error) (A1, A2) {
	return func(errs ...*error) (A1, A2) {
		CommonCheck(errs)(a1, a2)
		return a1, a2
	}
}

func Check2[A1, A2, A3 any](a1 A1, a2 A2, a3 A3) func(errs ...*error) (A1, A2, A3) {
	return func(errs ...*error) (A1, A2, A3) {
		CommonCheck(errs)(a1, a2, a3)
		return a1, a2, a3
	}
}

func Check3[A1, A2, A3, A4 any](a1 A1, a2 A2, a3 A3, a4 A4) func(errs ...*error) (A1, A2, A3, A4) {
	return func(errs ...*error) (A1, A2, A3, A4) {
		CommonCheck(errs)(a1, a2, a3, a4)
		return a1, a2, a3, a4
	}
}

func Check4[A1, A2, A3, A4, A5 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) func(errs ...*error) (A1, A2, A3, A4, A5) {
	return func(errs ...*error) (A1, A2, A3, A4, A5) {
		CommonCheck(errs)(a1, a2, a3, a4, a5)
		return a1, a2, a3, a4, a5
	}
}

func Check5[A1, A2, A3, A4, A5, A6 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) func(errs ...*error) (A1, A2, A3, A4, A5, A6) {
	return func(errs ...*error) (A1, A2, A3, A4, A5, A6) {
		CommonCheck(errs)(a1, a2, a3, a4, a5, a6)
		return a1, a2, a3, a4, a5, a6
	}
}

func Check6[A1, A2, A3, A4, A5, A6, A7 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) func(errs ...*error) (A1, A2, A3, A4, A5, A6, A7) {
	return func(errs ...*error) (A1, A2, A3, A4, A5, A6, A7) {
		CommonCheck(errs)(a1, a2, a3, a4, a5, a6, a7)
		return a1, a2, a3, a4, a5, a6, a7
	}
}

func Check7[A1, A2, A3, A4, A5, A6, A7, A8 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) func(errs ...*error) (A1, A2, A3, A4, A5, A6, A7, A8) {
	return func(errs ...*error) (A1, A2, A3, A4, A5, A6, A7, A8) {
		CommonCheck(errs)(a1, a2, a3, a4, a5, a6, a7, a8)
		return a1, a2, a3, a4, a5, a6, a7, a8
	}
}

func Check8[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) func(errs ...*error) (A1, A2, A3, A4, A5, A6, A7, A8, A9) {
	return func(errs ...*error) (A1, A2, A3, A4, A5, A6, A7, A8, A9) {
		CommonCheck(errs)(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return a1, a2, a3, a4, a5, a6, a7, a8, a9
	}
}

func Check9[A0, A1, A2, A3, A4, A5, A6, A7, A8, A9 any](a0 A0, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) func(errs ...*error) (A0, A1, A2, A3, A4, A5, A6, A7, A8, A9) {
	return func(errs ...*error) (A0, A1, A2, A3, A4, A5, A6, A7, A8, A9) {
		CommonCheck(errs)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return a0, a1, a2, a3, a4, a5, a6, a7, a8, a9
	}
}
