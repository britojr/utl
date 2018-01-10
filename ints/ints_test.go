package ints

import (
	"reflect"
	"sort"
	"testing"
)

func TestDifference(t *testing.T) {
	cases := []struct {
		a, b, want []int
	}{
		{[]int(nil), []int(nil), []int(nil)},
		{[]int{2, 3}, []int{}, []int{2, 3}},
		{[]int{}, []int{3, 4}, []int(nil)},
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6, 8}, []int{1, 3, 5}},
	}
	for _, tt := range cases {
		got := Difference(tt.a, tt.b)
		sort.Ints(got)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("(%v) - (%v) = (%v) !=(%v)", tt.a, tt.b, tt.want, got)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		is   []int
		want int
	}{
		{[]int{2, 3}, 3},
		{[]int{4, 5, 6, 1, 2, 3}, 6},
		{[]int{4, 5, -6, 1, 2, 3}, 5},
	}
	for _, tt := range cases {
		got := Max(tt.is)
		if tt.want != got {
			t.Errorf("max(%v) = (%v) != (%v)", tt.is, tt.want, got)
		}
	}
}
