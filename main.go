package main

import (
    "fmt"
    "os"
    "bufio"
    "./interpreter/lexer"
    "./interpreter/parser"
    "./scheme/number"
    "./scheme/boolean"
    "./scheme/symbol"
    //"./scheme/pair"
    "./scheme/emptylist"
)

func main() {
    stdin := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("goliath> ")
        lexer := lexer.New(stdin)
        parser := parser.New(lexer)
        sexp, err :=  parser.Parse()

        if err != nil {
            fmt.Println("Error")
            continue
        }

        switch sexp := sexp.(type) {
            case *number.Number:
                fmt.Println((*sexp).String())
            case *boolean.Boolean:
                fmt.Println((*sexp).String())
            case *symbol.Symbol:
                fmt.Println((*sexp).String())
            case *emptylist.EmptyList:
                fmt.Println((*sexp).String())

            default:
                fmt.Println("Unknown")
        }
    }

}
