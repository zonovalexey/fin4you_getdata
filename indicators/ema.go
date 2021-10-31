package indicators

import (
	"getdata/datastruct"
)

func Ema(data *[]datastruct.Tickerdata) {
	period := []int{5, 20, 65, 99, 140, 200}
	var sum, ema_start, a float64

	for _, per := range period {
		sum, ema_start = 0, 0
		a = 2 / (float64(per) + 1)
		for i := range *data {
			if i < per {
				sum += (*data)[i].Close
				switch per {
				case 5:
					(*data)[i].Ema5 = 0
				case 20:
					(*data)[i].Ema20 = 0
				case 65:
					(*data)[i].Ema65 = 0
				case 99:
					(*data)[i].Ema99 = 0
				case 140:
					(*data)[i].Ema140 = 0
				case 200:
					(*data)[i].Ema200 = 0
				}
			}
			if i == per {
				ema_start = sum / float64(per)
				switch per {
				case 5:
					(*data)[i].Ema5 = ema_start
				case 20:
					(*data)[i].Ema20 = ema_start
				case 65:
					(*data)[i].Ema65 = ema_start
				case 99:
					(*data)[i].Ema99 = ema_start
				case 140:
					(*data)[i].Ema140 = ema_start
				case 200:
					(*data)[i].Ema200 = ema_start
				}
			}
			if i > per {
				switch per {
				case 5:
					(*data)[i].Ema5 = a*(*data)[i].Close + (1-a)*ema_start
				case 20:
					(*data)[i].Ema20 = a*(*data)[i].Close + (1-a)*ema_start
				case 65:
					(*data)[i].Ema65 = a*(*data)[i].Close + (1-a)*ema_start
				case 99:
					(*data)[i].Ema99 = a*(*data)[i].Close + (1-a)*ema_start
				case 140:
					(*data)[i].Ema140 = a*(*data)[i].Close + (1-a)*ema_start
				case 200:
					(*data)[i].Ema200 = a*(*data)[i].Close + (1-a)*ema_start
				}
				ema_start = a*(*data)[i].Close + (1-a)*ema_start
			}
		}
	}
	// return data
}
