package grokking

func SortSelection(in []int) (out []int, err error) {
	whichLargest := func(values []int) int {
		var idx int
		largest := values[idx]
		for i, v := range values[1:] {
			if v > largest {
				idx = i + 1
				largest = v
			}
		}
		return idx
	}

	for len(in) > 0 {
		idx := whichLargest(in)
		out = append(out, in[idx])

		upper := idx + 1
		if upper > len(in) {
			upper = len(in)
		}
		in = append(in[0:idx], in[upper:]...)
	}
	return out, nil
}
