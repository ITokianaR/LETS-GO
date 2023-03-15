package main
import (
  "fmt"
)

func main() {
  // to have a better understading of pointer
  var name string = "John" // first we declare a string variable
  var ptr *string // then we create a pointer variable to store the memory adress

  ptr = &name // we assign the memory variable of "John" to the pointer variable

  fmt.Println("The memory adress of", name, "is", ptr)
}
