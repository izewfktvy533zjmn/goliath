package main

import (
    "fmt"
    "os"
    "bufio"
    "./interpreter/lexer"
    "./interpreter/parser"
    "./scheme"
    "./scheme/frame"
    "./scheme/evaluator"
)

func main() {
    stdin := bufio.NewScanner(os.Stdin)
    lexer := lexer.New(stdin)
    parser := parser.New(lexer)
    frame := frame.New(nil)
    evaluator := evaluator.New(frame)

    for {
        fmt.Print("goliath> ")
        if sexp, err :=  parser.Parse(); err != nil {
            fmt.Println(err)
            continue
        } else {
            if evaluatedSExp, err := evaluator.Evaluate(sexp); err != nil {
                fmt.Println(err)
            } else {
                fmt.Println(evaluatedSExp.(scheme.SExp).ToString())
            }
        }
    }

}
