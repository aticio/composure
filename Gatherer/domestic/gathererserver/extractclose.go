package gathererserver

func extractClose(klines []kline) []float64 {
	close := []float64{}

	for _, k := range klines {
		close = append(close, k.close)
	}

	return close
}
