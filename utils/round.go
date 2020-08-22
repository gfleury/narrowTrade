package utils

import "math"

func RoundUp(x float64, decimals int) float64 {
	roundDec := math.Pow(10, float64(decimals))
	return math.Ceil(x*roundDec) / roundDec
}

func Round(x float64, decimals int) float64 {
	roundDec := math.Pow(10, float64(decimals))
	return math.Round(x*roundDec) / roundDec
}

func RoundDown(x float64, decimals int) float64 {
	roundDec := math.Pow(10, float64(decimals))
	return math.Floor(x*roundDec) / roundDec
}
