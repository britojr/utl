// Package floats provides general functions for float64s
package floats

import "math"

const epslon = 1e-14 // default tolerace

// AlmostEqual compares two floats with a tolerance
func AlmostEqual(a, b float64, tolerance ...float64) bool {
	eps := epslon
	if len(tolerance) > 0 {
		eps = tolerance[0]
	}
	return math.Abs(a-b) < eps
}

// Max returns the max number
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		panic("floats: empty slice")
	}
	v := xs[0]
	for _, x := range xs {
		if x > v {
			v = x
		}
	}
	return v
}

// Min returns the min number
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		panic("floats: empty slice")
	}
	v := xs[0]
	for _, x := range xs {
		if x < v {
			v = x
		}
	}
	return v
}

// Sum returns the sum of all values
func Sum(fs []float64) (sum float64) {
	for _, v := range fs {
		sum += v
	}
	return
}

// Round float to a number of decimal places
func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
