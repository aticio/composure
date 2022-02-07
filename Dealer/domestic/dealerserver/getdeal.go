package dealerserver

import "fmt"

func getDeal(bi BulkInfo) {
	d := Deal{}
	if bi.LinRegSlope > 0 {
		d.Side = "BUY"
	} else {
		d.Side = "SELL"
	}
	fmt.Println(d)
}
