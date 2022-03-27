package genemap

func MapInt(list []int, cb func(int) int) []int {
	for i, itm := range list {
		list[i] = cb(itm)
	}

	return list
}
