// Package conv provides conversion functions
package conv

import (
	"fmt"
	"strconv"

	"github.com/britojr/kbn/utl/errchk"
)

// Atoi converts string to int
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	errchk.Check(err, fmt.Sprintf("Can't convert %v to int", s))
	return i
}

// Atof converts string to float64
func Atof(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	errchk.Check(err, fmt.Sprintf("Can't convert %v to float64", s))
	return i
}

// Satoi creates an int slice from a string slice
func Satoi(ss []string) []int {
	arr := make([]int, len(ss))
	for k, v := range ss {
		arr[k] = Atoi(v)
	}
	return arr
}

// Satof creates a float64 slice from a string slice
func Satof(ss []string) []float64 {
	arr := make([]float64, len(ss))
	for k, v := range ss {
		arr[k] = Atof(v)
	}
	return arr
}

// Sitou creates an uint64 array from an int array
func Sitou(is []int) []uint64 {
	arr := make([]uint64, len(is))
	for i, v := range is {
		arr[i] = uint64(v)
	}
	return arr
}

// Sitof creates an float64 array from an int array
func Sitof(is []int) []float64 {
	arr := make([]float64, len(is))
	for i, v := range is {
		arr[i] = float64(v)
	}
	return arr
}
