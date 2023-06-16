package err4

func Handle(errs ...*error) func(fn func()) {
	return func(fn func()) {
		for _, err := range errs {
			if err != nil && *err != nil {
				fn()
				break
			}
		}
	}
}
