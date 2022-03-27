//go:build generics

package geneplus

import (
	"github.com/MistaTwista/generigo/internal/util"
)

func addGenerics[T util.Addeable](a, b T) T {
	return a + b
}

func addGenericsVari[T util.Addeable](list ...T) T {
	var res T
	for _, itm := range list {
		res += itm
	}

	return res
}
