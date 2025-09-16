package lox

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must0(err error) {
	must(err)
}

func Must1[T any](val T, err error) T {
	must(err)
	return val
}

func Must2[T1, T2 any](val1 T1, val2 T2, err error) (T1, T2) {
	must(err)
	return val1, val2
}
