package slice

// ２つのスライスの和集合
// e.g. [1, 2, 4] + [1, 4, 5] => [1, 2, 4, 5]
func CulcUnion(s1, s2 []int64) []int64 {
	m := make(map[int64]struct{}, len(s1))

	for _, el := range s1 {
		m[el] = struct{}{}
	}

	for _, el := range s2 {
		m[el] = struct{}{}
	}

	r := make([]int64, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

// 2つのスライスの差集合
// e.g. [1, 2, 4] - [1, 4, 5] => [2]
func CulcMinus(s1, s2 []int64) []int64 {
	m := make(map[int64]struct{}, len(s1))

	for _, el := range s1 {
		m[el] = struct{}{}
	}

	for _, el := range s2 {
		delete(m, el)
	}

	r := make([]int64, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

// 2つのスライスの積集合
// e.g. [1, 2, 4] * [1, 4, 5] => [1, 4]
func CulcIntersection(s1, s2 []int64) []int64 {
	m := make(map[int64]struct{}, len(s1))

	for _, el := range s1 {
		m[el] = struct{}{}
	}

	r := make([]int64, 0)
	for _, el := range s2 {
		if _, ok := m[el]; !ok {
			continue
		}
		r = append(r, el)
	}

	return r
}
