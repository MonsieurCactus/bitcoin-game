package main

import "github.com/go-martini/martini"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "strconv"

type Price struct {
  Fifteen float64
  Last float64
  Buy float64
  Sell float64
  Symbol string
}

type Ticker struct
{
  USD  Price
  ISK  Price
  HKD  Price
  TWD  Price
  CHF  Price
  EUR  Price
  DKK  Price
  CLP  Price
  CAD  Price
  CNY  Price
  THB  Price
  AUD  Price
  SGD  Price
  KRW  Price
  JPY  Price
  PLN  Price
  GBP  Price
  SEK  Price
  NZD  Price
  BRL  Price
  RUB  Price
}

func main() {


  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/:cur", func(params martini.Params) string {
    resp, err := http.Get("https://blockchain.info/ticker")
    robots, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    	// handle error
    }
    defer resp.Body.Close()

    var t  map[string]Price

    err = json.Unmarshal(robots, &t)
    if err != nil {
    	// handle error
    }
    //return string(robots) + " " + strconv.FormatFloat(t.EUR.Buy, 'E', -1, 64)
    return string(params["cur"]) + ": " + strconv.FormatFloat(t[params["cur"]].Buy, 'f', 2, 64)
  })
  m.Run()
}
