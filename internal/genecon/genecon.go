package genecon

type Stringer interface {
	String() string
}

func GStringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}

	return ret
}

func Stringify(s []Stringer) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}

	return ret
}
