package main

import (
    "fmt"
    "os"
    "bufio"
    "./interpreter/lexer"
    "./interpreter/parser"
    "./scheme/number"
    "./scheme/boolean"
    "./scheme/identifier"
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

        switch sexp.(type) {
            case *number.Number:
                fmt.Println((*(sexp.(*number.Number))).String())

            case *boolean.Boolean:
                fmt.Println((*(sexp.(*boolean.Boolean))).String())

            case *identifier.Identifier:
                fmt.Println((*(sexp.(*identifier.Identifier))).String())

            case *emptylist.EmptyList:
                fmt.Println((*(sexp.(*emptylist.EmptyList))).String())

            default:
                fmt.Println("Unknown")
        }
    }

}
