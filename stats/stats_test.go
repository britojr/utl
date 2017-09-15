package stats

import (
	"reflect"
	"testing"

	"github.com/britojr/kbn/utl/floats"
)

func TestMean(t *testing.T) {
	cases := []struct {
		xs   []float64
		mean float64
	}{
		{[]float64{2, 2, 2}, 2},
		{[]float64{1, 2, 3}, 2},
		{[]float64{5, 4, 1, 2, 3, 6}, 3.5},
		{[]float64{12, 12, 12, 12, 13013}, 2612.2},
	}
	for _, tt := range cases {
		got := Mean(tt.xs)
		if !floats.AlmostEqual(tt.mean, got) {
			t.Errorf("wrong value,  want %v, got %v", tt.mean, got)
		}
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		xs   []float64
		mean float64
	}{
		{[]float64{1, 2, 3}, 2},
		{[]float64{2, 2, 2}, 2},
		{[]float64{5, 4, 1, 2, 3, 6}, 3.5},
		{[]float64{3, 1, 7}, 3},
	}
	for _, tt := range cases {
		got := Median(tt.xs)
		if !floats.AlmostEqual(tt.mean, got) {
			t.Errorf("wrong value,  want %v, got %v", tt.mean, got)
		}
	}
}

func TestMode(t *testing.T) {
	cases := []struct {
		xs   []float64
		want float64
	}{
		{[]float64{1, 2, 3}, 1},
		{[]float64{2, 2, 2, 1, 3, 3}, 2},
		{[]float64{5, 4, 1, 2, 3, 6}, 5},
		{[]float64{3, 1, 7, 1, 7, 7, 3}, 7},
	}
	for _, tt := range cases {
		got := Mode(tt.xs)
		if !floats.AlmostEqual(tt.want, got) {
			t.Errorf("mode %v != %v", tt.want, got)
		}
	}
}

func TestVariance(t *testing.T) {
	cases := []struct {
		xs   []float64
		want float64
	}{
		{[]float64{2, 2, 2}, 0},
		// {[]float64{1, 2, 3}, 2.0 / 3.0},
		// {[]float64{5, 4, 1, 2, 3, 6}, 2.916666667},
		// {[]float64{12, 12, 12, 12, 13013}, 27044160.16},
	}
	for _, tt := range cases {
		got := Variance(tt.xs)
		if !floats.AlmostEqual(tt.want, got, 1e-6) {
			t.Errorf("wrong value,  want %v, got %v", tt.want, got)
		}
	}
}

func TestStdev(t *testing.T) {
	cases := []struct {
		xs []float64
		sd float64
	}{
		{[]float64{2, 2, 2}, 0},
		// {[]float64{1, 2, 3}, 0.816496581},
		// {[]float64{5, 4, 1, 2, 3, 6}, 1.707825128},
		// {[]float64{12, 12, 12, 12, 13013}, math.Sqrt(27044160.16)},
	}
	for _, tt := range cases {
		got := Stdev(tt.xs)
		if !floats.AlmostEqual(tt.sd, got, 1e-6) {
			t.Errorf("wrong value,  want %v, got %v", tt.sd, got)
		}
	}
}

func TestDirichlet(t *testing.T) {
	cases := []struct {
		alpha float64
		l     int
	}{
		{3.2, 4},
		{0.1, 8},
		{0.01, 2},
		{5, 1},
		{1.5, 9},
	}
	for _, tt := range cases {
		values := make([]float64, tt.l)
		Dirichlet1(tt.alpha, values)
		if !floats.AlmostEqual(1, floats.Sum(values)) {
			t.Errorf("not normalized %v", values)
		}
	}

	// test different outcomes
	alpha, l := 0.7, 8
	a, b := make([]float64, l), make([]float64, l)
	Dirichlet1(alpha, a)
	Dirichlet1(alpha, b)
	count := 0
	for i := 0; i < l; i++ {
		if floats.AlmostEqual(a[i], b[i]) {
			count++
		}
	}
	if count == l {
		t.Errorf("Sampled the same distribution:\n%v\n%v", a, b)
	}
}

func TestSetUniform(t *testing.T) {
	cases := []struct {
		xs, want []float64
	}{{
		[]float64{10, 20, 30, 40},
		[]float64{0.25, 0.25, 0.25, 0.25},
	}, {
		[]float64{2, 2, 2, 2, 2, 2, 2, 2},
		[]float64{0.125, 0.125, 0.125, 0.125, 0.125, 0.125, 0.125, 0.125},
	}, {
		[]float64{},
		[]float64{},
	}}
	for _, tt := range cases {
		Uniform(tt.xs)
		got := tt.xs
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("uniform: %v != %v", tt.want, got)
		}
	}
}

func TestSetRandom(t *testing.T) {
	cases := []struct {
		xs []float64
	}{{
		[]float64{10, 20, 30, 40},
	}, {
		[]float64{5.5},
	}, {
		[]float64{0, 0, 0},
	}, {
		[]float64{3, 4},
	}}
	for _, tt := range cases {
		Random(tt.xs)
		got := tt.xs
		if !floats.AlmostEqual(1, floats.Sum(got)) {
			t.Errorf("not normalized, sum %v", floats.Sum(got))
		}
	}
	// test different outcomes
	xs := []float64{1, 2, 3}
	Random(xs)
	values := append([]float64(nil), xs...)
	Random(xs)
	count := 0
	for i := range values {
		if floats.AlmostEqual(values[i], xs[i]) {
			count++
		}
	}
	if count == len(values) {
		t.Errorf("Sampled the same distribution:\n%v\n%v", values, xs)
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		values, normalized []float64
	}{{
		[]float64{0.15, 0.25, 0.35, 0.25},
		[]float64{0.15, 0.25, 0.35, 0.25},
	}, {
		[]float64{15, 25, 35, 25},
		[]float64{0.15, 0.25, 0.35, 0.25},
	}, {
		[]float64{10, 20, 30, 40, 50, 60, 70, 80},
		[]float64{1.0 / 36, 2.0 / 36, 3.0 / 36, 4.0 / 36, 5.0 / 36, 6.0 / 36, 7.0 / 36, 8.0 / 36},
	}, {
		[]float64{0.15},
		[]float64{1},
	}}
	for _, tt := range cases {
		Normalize(tt.values)
		if !reflect.DeepEqual(tt.values, tt.normalized) {
			t.Errorf("want %v, got %v", tt.normalized, tt.values)
		}
	}
}
