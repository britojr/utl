package conv

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"1", 1},
		{"+0", 0},
		{"-2", -2},
	}
	for _, tt := range cases {
		got := Atoi(tt.s)
		if tt.want != got {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestAtof(t *testing.T) {
	cases := []struct {
		s    string
		want float64
	}{
		{"1", 1},
		{"-0.00005", -0.00005},
		{"+1.2", 1.2},
		{"3.14", 3.14},
	}
	for _, tt := range cases {
		got := Atof(tt.s)
		if tt.want != got {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestAtob(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"true", true},
		{"True", true},
		{"false", false},
		{"False", false},
	}
	for _, tt := range cases {
		got := Atob(tt.s)
		if tt.want != got {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestSatoi(t *testing.T) {
	cases := []struct {
		s    []string
		want []int
	}{
		{[]string{"1"}, []int{1}},
		{[]string{"-1", "7", "-9", "3"}, []int{-1, 7, -9, 3}},
	}
	for _, tt := range cases {
		got := Satoi(tt.s)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestSatof(t *testing.T) {
	cases := []struct {
		s    []string
		want []float64
	}{
		{[]string{"1"}, []float64{1}},
		{[]string{"-1.7", "3.17", "-0.00019", "3"}, []float64{-1.7, 3.17, -0.00019, 3}},
	}
	for _, tt := range cases {
		got := Satof(tt.s)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestSftoa(t *testing.T) {
	cases := []struct {
		fs   []float64
		want []string
	}{
		{[]float64{1}, []string{strconv.FormatFloat(1, 'E', -1, 64)}},
		{[]float64{-1.7, 3.17, -0.00019, 3}, []string{
			strconv.FormatFloat(-1.7, 'E', -1, 64),
			strconv.FormatFloat(3.17, 'E', -1, 64),
			strconv.FormatFloat(-0.00019, 'E', -1, 64),
			strconv.FormatFloat(3, 'E', -1, 64),
		}},
	}
	for _, tt := range cases {
		got := Sftoa(tt.fs)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestSitou(t *testing.T) {
	cases := []struct {
		s    []int
		want []uint64
	}{
		{[]int{1}, []uint64{1}},
		{[]int{1, 7, 0, 3}, []uint64{1, 7, 0, 3}},
	}
	for _, tt := range cases {
		got := Sitou(tt.s)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}

func TestSitof(t *testing.T) {
	cases := []struct {
		s    []int
		want []float64
	}{
		{[]int{1}, []float64{1}},
		{[]int{-1, 7, -9, 3}, []float64{-1, 7, -9, 3}},
	}
	for _, tt := range cases {
		got := Sitof(tt.s)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("wrong conversion %v != %v", tt.want, got)
		}
	}
}
