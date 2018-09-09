package main

import "fmt"

type Vertex struct {
  Lat, Lon float64
}

var m = map[string]Vertex{
  "Bell Labs": {40.68433, -74.3996},
  "Google": {40.68433, -74.3996},
}

func main()  {
  fmt.Println(m)
}