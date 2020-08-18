package utils

import "math"

func RoundUp(x float64) float64 {
	return math.Ceil(x*100) / 100
}

func Round(x float64) float64 {
	return math.Round(x*100) / 100
}

func RoundDown(x float64) float64 {
	return math.Floor(x*100) / 100
}
