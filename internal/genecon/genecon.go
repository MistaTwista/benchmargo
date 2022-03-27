//go:build generics

package genecon

import "github.com/MistaTwista/benchmargo/internal/util"

func GStringify[T util.Stringer](s ...T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}

	return ret
}
