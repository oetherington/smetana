package smetana

// A helper function to create an "id" attribute.
func Id(id string) Attr {
	return Attr{"id", id}
}

// Merge the `src` map into the `dst` map in place, replacing any duplicate
// keys. `src` is unchanged.
func MergeMaps[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

type ordered interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		string
}

func min[T ordered](a T, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func max[T ordered](a T, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func clamp[T ordered](value T, min T, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Transform an array using a given function. This is equivalent to `map` in
// Haskell or Javascript. For example:
//
//	Xform([]int32{1, -2, 3}, func(a int32) uint32 {
//		if a < 0 {
//			a = -a
//	 	}
//	 	return uint32(a)
//	 })
//
// will return []uint32{1, 2, 3}.
func Xform[A any, B any](values []A, xform func(a A) B) []B {
	out := make([]B, len(values))
	for i, value := range values {
		out[i] = xform(value)
	}
	return out
}
