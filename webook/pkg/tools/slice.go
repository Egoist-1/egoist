package tools

func Map[Src any, Dst any](src []Src, fn func(idx int, s Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for idx, s := range src {
		dst[idx] = fn(idx, s)
	}
	return dst
}
