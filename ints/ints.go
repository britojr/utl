package ints

// Difference returns the difference of two slices:
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

// Max returns the maximum value in slice
// if the slice is empty, Max will panic
func Max(is []int) int {
	if len(is) == 0 {
		panic("ints: slice with zero length")
	}
	max := is[0]
	for _, v := range is {
		if v > max {
			max = v
		}
	}
	return max
}
