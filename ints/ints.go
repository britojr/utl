package ints

// Difference retirns the difference of two slices:
//		c = a-b
func Difference(a, b []int) (c []int) {
	m := make(map[int]struct{})
	for _, u := range b {
		m[u] = struct{}{}
	}
	for _, u := range a {
		if _, ok := m[u]; !ok {
			c = append(c, u)
		}
	}
	return
}
