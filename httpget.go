package main

import (
  "fmt"
  "net/http"
  "time"
  "sync"
)

var array = []string{
  "aceon",
  "activeoxy",
  "aerospan",
  "ajmalicine",
  "alfentanil",
  "ammonium",
  "ampicillin",
  "anesgerm",
  "anistreplase",
  "betavet",
  "bio-mycin",
  "brovana",
  "camphor",
  "capsin",
  "centany",
  "cipro",
  "cortisone",
  "cyproheptadine",
  "dehydroabietate",
  "propane",
  "proparacaine",
  "promethegan",
  "premarin",
  "patanol",
  "molybdenum",
  "methenolone",
  "menthol",
  "melamin",
  "magtrate",
  "lenograstim",
  "hexoprenaline",
}


func main() {
  var wg sync.WaitGroup
  timestart := time.Now()
  wg.Add(len(array))
  ch := make(chan string)
  for i:=0 ; i < len(array) ; i++ {
    url := array[i]
    fmt.Println("fetching ", url)
    go func(url1 string){
      defer wg.Done()
      resp, _ := http.Get(url1)
      ch <- resp.Status
    }("http://rxnav.nlm.nih.gov/REST/drugs?name=" + url)
  }

  go func() {
        for response := range ch {
            fmt.Println(response)
        }
  }()

  wg.Wait()

  timeend := time.Now()
  fmt.Println(timeend.Sub(timestart))
}