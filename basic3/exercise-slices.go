package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  tmp := make([][]uint8, dy)
  for i := range tmp {
    tmp[i] = make([]uint8, dx)
  }
  for y := 0; y < dy; y++ {
    for x :=0; x < dx; x++ {
      tmp[y][x] = uint8((x + y) / 2)
    }
  }
  return tmp
}

func main() {
  pic.Show(Pic)
}