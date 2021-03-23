package main

import "fmt"

//importing the quote module from pkg.go.dev
//go mod tidy downloaded the rsc.io/quote module
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}