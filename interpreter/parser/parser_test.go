package parser

import (
    "os"
    "fmt"
    "bufio"
    "errors"
    "testing"
    "../lexer"
    "../../scheme"
)

func TestParse(test *testing.T) {
    inputText := "10"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Close()

    fp, err = os.Open("test.scm")
    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    l := lexer.New(in)
    parser := New(l)

    expect := *(&scheme.Number{Num: 10})
    tmp, err := parser.Parse()

    if err != nil {
        test.Errorf("%s", err)
    }

    actual := *(tmp.(*scheme.Number))

    if expect != actual {
        test.Errorf("%d != %d", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestParse_error(test *testing.T) {
    inputText := "test"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Close()

    fp, err = os.Open("test.scm")
    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    l := lexer.New(in)
    parser := New(l)

    expect := errors.New("ParseErrorException").Error()
    tmp , err := parser.Parse()
    actual := err.Error()

    if tmp != nil || err == nil {
        test.Errorf("%s", err)
    }

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}
