package main
import (
  "fmt"
)

func sum(x, y int) int {
  return x + y
}

func substract(x, y int) int {
  return x - y
}

func divide(x, y int) int {
  return x / y
}

func multiply(x, y int) int {
  return x * y
}

//just a simple calculator to better learn switch case
func main() {
  var num1 int
  var num2 int
  var op string

  fmt.Print("Enter a number :")
  fmt.Scan(&num1)

  fmt.Print("Choose an operator + / - *")
  fmt.Scan(&op)

  fmt.Print("Enter a number :")
  fmt.Scan(&num2)

  switch op {
    case "+" :
      sum := sum(num1, num2)
      fmt.Printf("%d + %d = %d", num1, num2, sum)

    case "-" :
      sub := substract(num1, num2)
      fmt.Printf("%d - %d = %d", num1, num2, sub)

    case "/" :
      div := divide(num1, num2)
      fmt.Printf("%d / %d = %d", num1, num2, div)

    case "*" :
      mul := multiply(num1, num2)
      fmt.Printf("%d * %d = %d", num1, num2, mul)
  }
}
