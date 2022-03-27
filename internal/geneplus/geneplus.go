//go:build generics

package geneplus

import (
	"github.com/MistaTwista/benchmargo/internal/util"
)

func addGenerics[T util.Addable](a, b T) T {
	return a + b
}

func AddGenericsVari[T util.Addable](list ...T) T {
	var res T
	for _, itm := range list {
		res += itm
	}

	return res
}
