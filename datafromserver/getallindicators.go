package datafromserver

import (
	"fmt"
	"time"
)

func Getall() {
	data := server_get_api_struct("aapl", "1d")
	for _, v := range data {
		fmt.Println(time.Unix(v.date, 0), v.open, v.close, (v.close/v.open-1)*100)
	}
}
