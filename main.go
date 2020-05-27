package main

import (
    "fmt"
    "os"
    "bufio"
    "./interpreter/lexer"
    "./interpreter/parser"
    "./scheme/number"
    "./scheme/boolean"
)

func main() {
    stdin := bufio.NewScanner(os.Stdin)

    fmt.Print("goliath> ")
    lexer := lexer.New(stdin)
    parser := parser.New(lexer)
    sexp, _ :=  parser.Parse()

    switch sexp.(type) {
        case *number.Number:
            fmt.Println((*(sexp.(*number.Number))).String())

        case *boolean.Boolean:
            fmt.Println((*(sexp.(*boolean.Boolean))).String())

        default:
            fmt.Println("Unknown")
    }

}
