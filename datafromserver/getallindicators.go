package datafromserver

import (
	"encoding/json"
	"fmt"
	"getdata/indicators"
	"log"
	"net/http"
)

// var mu sync.Mutex

func Getall() {
	//http://localhost:8000/info?ticker=aapl&period=1d
	http.HandleFunc("/info", info)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func info(w http.ResponseWriter, r *http.Request) {
	// mu.Lock()
	query := r.URL.Query()
	ticker, ok1 := query["ticker"]
	period, ok2 := query["period"]
	per := "1d"
	if ok1 {
		if ok2 {
			per = period[0]
		}
		data := server_get_api_struct(ticker[0], per)
		indicators.Ema(&data)
		answer, err := json.Marshal(&data)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(answer)
		// data = nil
		// answer = nil
	}
	// mu.Unlock()
}
