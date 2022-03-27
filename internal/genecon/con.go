package genecon

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"strconv"
)

type Str struct {
	S string
}

func (s Str) String() string {
	return s.S
}

var res []string

func gena(n int) []Str {
	res := make([]Str, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, Str{S: strconv.Itoa(i)})
	}
	return res
}

func NewStringer(list []Str) []util.Stringer {
	res := make([]util.Stringer, 0, len(list))
	for _, itm := range list {
		res = append(res, util.Stringer(itm))
	}

	return res
}

func Stringify(list ...util.Stringer) (ret []string) {
	for _, v := range list {
		ret = append(ret, v.String())
	}

	return ret
}
