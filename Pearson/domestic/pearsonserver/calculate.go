package pearsonserver

import (
	"github.com/dgryski/go-onlinestats"
)

func calculate(p price) float64 {
	var x []float64
	for i := range p.Close {
		x = append(x, float64(i))
	}

	pr := onlinestats.Pearson(x, p.Close)
	return pr
}
