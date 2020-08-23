package utils

import (
	"math"
)

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

func RoundNearestTick(price, percentage, tickSize float64) float64 {
	return Round(math.Round(((price-(price*percentage)/100)/tickSize))*tickSize, GetTickDecimals(tickSize))
}

func GetTickDecimals(tickSize float64) int {
	if tickSize > 1 {
		return 0
	}
	count := 1
	for ; math.Pow(10, float64(count))*tickSize < 1; count++ {
	}
	return count
}
