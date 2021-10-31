package datafromserver

import (
	"fmt"
	"getdata/indicators"
	"math"
	"time"
)

func Getall() {
	data := server_get_api_struct("aapl", "1d")
	indicators.Ema(&data)
	for _, v := range data {
		fmt.Printf("Date: %v, open: %v, close: %v, percent: %v, ema5: %v, ema20 %v\n", time.Unix(v.Date, 0), v.Open, v.Close, math.Round((v.Close/v.Open-1)*100*100)/100, v.Ema5, v.Ema20)
	}
}
