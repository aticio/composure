package slopeserver

import (
	"github.com/GaryBoone/GoStats/stats"
)

func calculate(p Price) float64 {
	var x []float64
	for i := range p.Close {
		x = append(x, float64(i))
	}

	slope, _, _, _, _, _ := stats.LinearRegression(x, p.Close)
	return slope
}
