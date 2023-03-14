package main

import (
  "fmt"
)

func sum(x, y int) int {
  return x + y
}

func main() {
  nb := 5
  nbr := 10
  fmt.Println(nb, "+", nbr, "=", sum(nb, nbr)) 
}
