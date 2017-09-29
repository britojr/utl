package stats

import (
	"math"

	gfloats "gonum.org/v1/gonum/floats"
)

// MSE calculates mean squared error
func MSE(xs, ys []float64) (mse float64) {
	for i, v := range xs {
		mse += (v - ys[i]) * (v - ys[i])
	}
	return mse / float64(len(xs))
}

// CrossEntropy calculates cross entropy of two distributions
func CrossEntropy(xs, ys []float64) (c float64) {
	for i, v := range xs {
		// if ys[i] != 0 {
		if v != 0 {
			c -= v * math.Log(ys[i])
		}
	}
	return
}

// MaxAbsErr returns the maximum absolute error
func MaxAbsErr(xs, ys []float64) (e float64) {
	for i := range xs {
		x := math.Abs(xs[i] - ys[i])
		if e < x {
			e = x
		}
	}
	return
}

// HellDist calculates Hellinger distance
func HellDist(xs, ys []float64) (e float64) {
	for i := range xs {
		e += math.Pow(math.Sqrt(xs[i])-math.Sqrt(ys[i]), 2)
	}
	return e / math.Sqrt2
}

// LNormDiff calculates the l-norm of the difference of two arrays
func LNormDiff(xs, ys []float64, l float64) float64 {
	return gfloats.Distance(xs, ys, l)
}

// IntsLNormDiff calculates the l-norm of the difference of two int arrays
func IntsLNormDiff(xs, ys []int, l int) (d float64) {
	for i := range xs {
		d += math.Pow(float64(xs[i]-ys[i]), 2)
	}
	return math.Sqrt(d)
}

// MatMSE compares two float matrixes and return the Mean Squared Error
func MatMSE(e, a [][]float64) (mse float64) {
	for i := range e {
		mse += MSE(e[i], a[i])
	}
	return mse / float64(len(e))
}

// MatCrossEntropy compares two float matrixes and return the Mean cross entropy
func MatCrossEntropy(e, a [][]float64) (c float64) {
	for i := range e {
		c += CrossEntropy(e[i], a[i])
	}
	return c / float64(len(e))
}

// MatDistance compares two float matrixes and return the Mean L distance
func MatDistance(e, a [][]float64, l float64) (c float64) {
	for i := range e {
		c += gfloats.Distance(e[i], a[i], l)
	}
	return c / float64(len(e))
}

// MatMaxAbsErr compares two float matrixes and return the Mean Abs Error
func MatMaxAbsErr(e, a [][]float64) (c float64) {
	for i := range e {
		c += MaxAbsErr(e[i], a[i])
	}
	return c / float64(len(e))
}

// MatHellDist compares two float matrixes and return the Mean Hellinger distance
func MatHellDist(e, a [][]float64) (c float64) {
	for i := range e {
		c += HellDist(e[i], a[i])
	}
	return c / float64(len(e))
}
