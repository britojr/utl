package stats

import (
	"testing"

	"github.com/britojr/utl/floats"
)

func TestMSE(t *testing.T) {
	cases := []struct {
		xs, ys []float64
		mse    float64
	}{{
		[]float64{0.15, 0.25, 0.35, 0.25},
		[]float64{0.15, 0.25, 0.35, 0.25},
		0,
	}, {
		[]float64{0.15, 0.25, 0.35, 0.25},
		[]float64{0.10, 0.35, 0.34, 0.26},
		0.003175,
	}}
	for _, tt := range cases {
		got := MSE(tt.xs, tt.ys)
		if !floats.AlmostEqual(tt.mse, got) {
			t.Errorf("want %v, got %v", tt.mse, got)
		}
	}
}

func TestCrossEntropy(t *testing.T) {
	cases := []struct {
		xs, ys []float64
		d      float64
	}{{
		[]float64{0.15, 0.25, 0.60},
		[]float64{0.15, 0.25, 0.60},
		0.937636962,
	}, {
		[]float64{0.15, 0.25, 0.60},
		[]float64{0.10, 0.35, 0.55},
		0.966545496,
	}}
	for _, tt := range cases {
		got := CrossEntropy(tt.xs, tt.ys)
		if !floats.AlmostEqual(tt.d, got, 1e-9) {
			t.Errorf("want %v, got %v", tt.d, got)
		}
	}
}

func TestMaxAbsErr(t *testing.T) {
	cases := []struct {
		xs, ys []float64
		d      float64
	}{{
		[]float64{0.15, 0.25, 0.60},
		[]float64{0.15, 0.25, 0.60},
		0,
	}, {
		[]float64{0.15, 0.25, 0.60},
		[]float64{0.10, 0.35, 0.55},
		0.1,
	}}
	for _, tt := range cases {
		got := MaxAbsErr(tt.xs, tt.ys)
		if !floats.AlmostEqual(tt.d, got, 1e-9) {
			t.Errorf("want %v, got %v", tt.d, got)
		}
	}
}

func TestHellDist(t *testing.T) {
	cases := []struct {
		xs, ys []float64
		d      float64
	}{{
		[]float64{0.15, 0.25, 0.60},
		[]float64{0.15, 0.25, 0.60},
		0,
	}, {
		[]float64{0.15, 0.25, 0.60},
		[]float64{0.10, 0.35, 0.55},
		0.010274628,
	}}
	for _, tt := range cases {
		got := HellDist(tt.xs, tt.ys)
		if !floats.AlmostEqual(tt.d, got, 1e-9) {
			t.Errorf("want %v, got %v", tt.d, got)
		}
	}
}

func TestMatMSE(t *testing.T) {
	cases := []struct {
		e, a [][]float64
		d    float64
	}{{
		[][]float64{
			{.9, .1},
			{.1, .7, .2},
			{.8, .20001},
			{.5, .5},
		},
		[][]float64{
			{.9, .1},
			{.1, .7, .2},
			{.8, .20001},
			{.5, .5},
		},
		0,
	}, {
		[][]float64{
			{.9, .1},
			{.1, .7, .2},
		},
		[][]float64{
			{.8, .2},
			{.2, .7, .1},
		},
		0.008333333,
	}}
	for _, tt := range cases {
		got := MatMSE(tt.e, tt.a)
		if !floats.AlmostEqual(tt.d, got, 1e-5) {
			t.Errorf("want %v, got %v", tt.d, got)
		}
	}
}
