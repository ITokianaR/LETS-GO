package main
import (
  "fmt"
)

func main() {
  // to have a better understading of pointer
  var name string = "John"
  var ptr *string

  ptr = &name

  fmt.Println("The memory adress of", name, "is", ptr)
}
