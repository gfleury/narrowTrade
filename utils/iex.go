package utils

import (
	iex "github.com/goinvest/iexcloud/v2"
)

func IEXRecomendationReduce(rs []iex.Recommendation) *iex.Recommendation {
	ret := &iex.Recommendation{}

	split := 1
	for _, r := range rs {
		ret.BuyRatings = (r.BuyRatings + ret.BuyRatings) / split
		ret.ConsensusRating = (r.ConsensusRating + ret.ConsensusRating) / float64(split)
		ret.HoldRatings = (r.HoldRatings + ret.HoldRatings) / split
		ret.NoRatings = (r.NoRatings + ret.NoRatings) / split
		ret.OverweightRatings = (r.OverweightRatings + ret.OverweightRatings) / split
		ret.SellRatings = (r.SellRatings + ret.SellRatings) / split
		ret.UnderweightRatings = (r.UnderweightRatings + ret.UnderweightRatings) / split
		split = 2
	}

	return ret
}

func IEXBBandsReduction(bbands iex.Indicator, lastItems int) []float64 {
	ret := make([]float64, len(bbands.Indicator))
	for idx, indicator := range bbands.Indicator {
		for _, value := range indicator[len(indicator)-lastItems:] {
			ret[idx] += value
		}
		ret[idx] = ret[idx] / float64(lastItems)
	}

	return ret
}
