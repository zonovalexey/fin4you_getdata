package datafromserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func server_get_api_struct(tickername string, interval string) []tickerdata {

	var data yahoofinancestruct

	datestart := time.Now().Unix() - 50*365*24*60*60
	dateend := time.Now().Unix()

	var path = "https://query1.finance.yahoo.com/v8/finance/chart/" + tickername + "?period1=" + fmt.Sprint(datestart) + "&period2=" + fmt.Sprint(dateend) + "&interval=" + interval

	resp, err := http.Get(path)

	if err != nil {
		fmt.Println("Ошибка соединения с сервером при получении данных по тикеру \"" + tickername + "\"")
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка чтения данных по тикеру \"" + tickername + "\" с сервер")
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Ошибка преобразования данных задания \"" + tickername + "\" в JSON формат")
		log.Fatal(err)
	}

	var data_struct []tickerdata

	for i, res := range data.Chart.Result {
		if i == 0 {
			for k, quote := range res.Indicators.Quote {
				if k == 0 {
					for t := 0; t < len(quote.Open); t++ {
						data_struct = append(data_struct, tickerdata{date: res.Timestamp[t], open: quote.Open[t], close: quote.Close[t]})
					}
				}
			}
		}
	}

	return data_struct
}
