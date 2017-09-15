package stats

import (
	"testing"

	"github.com/britojr/utl/floats"
)

func TestGaussianKernel(t *testing.T) {
	cases := []struct {
		x  float64
		gk float64
	}{
		// {2, math.Pow(2*math.Pi, -.5) * math.Exp(-.5*(2*2))},
		{2, 0.05399096651318806},
		{7, 9.134720408364594e-12},
		{-7, 9.134720408364594e-12},
		{0.1, 0.3969525474770118},
		{-0.1, 0.3969525474770118},
		{0, 0.3989422804014327},
	}
	for _, tt := range cases {
		got := GaussianKernel(tt.x)
		if !floats.AlmostEqual(tt.gk, got) {
			t.Errorf("wrong value,  want %v, got %v", tt.gk, got)
		}
	}
}
