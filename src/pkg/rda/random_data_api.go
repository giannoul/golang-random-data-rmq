package rda 

import (
	"math/rand"
	"io"
	"fmt"
	"net/http"
	"encoding/json"
	"golang.org/x/time/rate"
	"time"
)

type OptFunc func(*Opts)

type Opts struct {
	Records int 
	ResponseType string
}

func defaultOpts() Opts {
	return Opts{
		Records: 1,
		ResponseType: "json",
	}
}

func GetMultipleRecords(opts *Opts) {
	var n int
	for n = rand.Intn(10); n < 1; n = rand.Intn(10) {
	}
	opts.Records = n
} 

type RandomAPI struct {
	Url string
	Opts Opts
}

func NewRandomAPI(url string, opts ...OptFunc) *RandomAPI {
	o := defaultOpts()
	for _,fn := range opts {
		fn(&o)
	}
	return &RandomAPI{
		Url: url,
		Opts: o,
	}
}

func (r RandomAPI) GetData() []byte {
	defaultResponse,_ := json.Marshal(struct{}{})
	url := fmt.Sprintf("%s?size=%d&response_type=%s",r.Url,r.Opts.Records,r.Opts.ResponseType)
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Println("RandomAPI.GetData http error", err)
		return defaultResponse
	}
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println("RandomAPI.GetData io.ReadAll error", err)
		return defaultResponse
	}
	return body
}

func (r RandomAPI) GetDataWithRateLoop(rps int, burst int ,ch chan []byte) {
    limiter := rate.NewLimiter(rate.Every(time.Duration(rps)*time.Second), burst) // max of "burst" requests and then "rps" more requests per second
	for ;; {
        if !limiter.Allow() {
			fmt.Println(r.Url, " - getting data")
            ch <- r.GetData()
        } else {
            fmt.Println(r.Url, " - rate limit reached, waiting")
        }
	}
}