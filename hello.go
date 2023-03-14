package main
import "fmt"

const (
        StatusOK                   = 200
        StatusCreated              = 201
        StatusAccepted             = 202
        StatusNonAuthoritativeInfo = 203
        StatusNoContent            = 204
        StatusResetContent         = 205
        StatusPartialContent       = 206
        Big   = 1 << 62
	      Small = Big >> 61
        Truth = false
)

func main() {
    name, location := "Itokiana", "Anjanahary"
    age := 19
    fmt.Printf("%s is a %d year-old boy from %s", name, age, location)
    fmt.Println(Truth)
}
