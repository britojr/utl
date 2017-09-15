package stats

import "math"

// GaussianKernel computes gaussian kernel function:
//		e^(-x^2 / 2) / sqrt(2*pi)
func GaussianKernel(x float64) float64 {
	return math.Pow(2*math.Pi, -.5) * math.Exp(-.5*(x*x))
}
