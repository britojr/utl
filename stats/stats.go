package stats

import (
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/britojr/kbn/utl/floats"
	gfloats "github.com/gonum/floats"
	"github.com/gonum/stat"
	"github.com/gonum/stat/distuv"
)

// Mean calculates the Mean of a float64 slice
func Mean(xs []float64) float64 {
	return floats.Sum(xs) / float64(len(xs))
}

// Median calculates the media of a float64 slice
func Median(xs []float64) (v float64) {
	aux := append([]float64(nil), xs...)
	sort.Float64s(aux)
	i := len(aux) / 2
	if len(aux)%2 != 0 {
		v = aux[i]
	} else {
		v = (aux[i] + aux[i-1]) / 2
	}
	return
}

// Mode calculates the mode of a float64 slice
func Mode(xs []float64) (v float64) {
	d := make(map[float64]int)
	c := 0
	for _, x := range xs {
		d[x]++
		if d[x] > c {
			c = d[x]
			v = x
		}
	}
	return
}

// Variance calculates the variance of a float64 slice
func Variance(xs []float64) (v float64) {
	return stat.Variance(xs, nil)
}

// Stdev calculates the standard deviation of a float64 slice
func Stdev(xs []float64) float64 {
	return math.Sqrt(Variance(xs))
}

// Dirichlet1 sample values as a Dirichlet distribution using single alpha parameter
func Dirichlet1(alpha float64, values []float64) {
	if alpha == 0 {
		panic("alpha != 0 needed for dirichlet dirtribution")
	}
	rndsrc := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := range values {
		values[i] = distuv.Gamma{Alpha: alpha, Beta: 1, Source: rndsrc}.Rand()
	}
	Normalize(values)
}

// Uniform sets values uniformly
func Uniform(values []float64) {
	n := float64(len(values))
	for i := range values {
		values[i] = 1.0 / n
	}
}

// Random sets random values
func Random(values []float64) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range values {
		values[i] = rand.Float64()
	}
	Normalize(values)
}

// Normalize normalizes the slice so all values sum to one
func Normalize(fs []float64) {
	sum := floats.Sum(fs)
	if sum == 0 {
		panic("trying to normalize a zero slice")
	}
	for i, v := range fs {
		fs[i] = v / sum
	}
}

// Distance functions

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
