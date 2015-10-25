package main

import "github.com/go-martini/martini"

import "net/http"


import "io/ioutil"



func main() {




  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/EUR", func() string {


    resp, err := http.Get("https://blockchain.info/ticker")
    robots, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    	// handle error
    }
    defer resp.Body.Close()

    return string(robots)
  })
  m.Run()
}
