package geneapply

func Apply(fns ...func(int) int) int {
	var r int
	for _, fn := range fns {
		r = fn(r)
	}

	return r
}

func Identity(n int) func(int) int {
	return func(_ int) int {
		return n
	}
}

func Multiply(i int) int {
	return i * i
}

func Twice(i int) int {
	return i + i
}

func MultiplyWith(n int) func(int) int {
	return func(i int) int {
		return i * n
	}
}
