package main

import "fmt"

//importing the quote module from pkg.go.dev
import "rsc.io/quote"


func main() {
    fmt.Println(quote.Go())
}