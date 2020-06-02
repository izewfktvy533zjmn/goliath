package main

import (
    "fmt"
    "os"
    "bufio"
    "./interpreter/lexer"
    "./interpreter/parser"
    "./scheme"
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
            case scheme.Scheme:
                fmt.Println(sexp.ToString())

            default:
                fmt.Println("Unknown")
        }
    }

}
