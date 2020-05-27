package parser

import (
    "os"
    "fmt"
    "bufio"
    "errors"
    "testing"
    "../lexer"
    "../../scheme/number"
    "../../scheme/boolean"
)

func TestParse_Number(test *testing.T) {
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

    expect := *(number.New(10))
    tmp, err := parser.Parse()

    if err != nil {
        test.Errorf("%s", err)
    }

    actual := *(tmp.(*number.Number))

    if expect != actual {
        test.Errorf("%d != %d", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestParse_Boolean_true(test *testing.T) {
    inputText := "#t"
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

    expect := *(boolean.New(true))
    tmp, err := parser.Parse()

    if err != nil {
        test.Errorf("%s", err)
    }

    actual := *(tmp.(*boolean.Boolean))

    if expect != actual {
        test.Errorf("%s != %s", expect.String(), actual.String())
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestParse_Boolean_false(test *testing.T) {
    inputText := "#f"
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

    expect := *(boolean.New(false))
    tmp, err := parser.Parse()

    if err != nil {
        test.Errorf("%s", err)
    }

    actual := *(tmp.(*boolean.Boolean))

    if expect != actual {
        test.Errorf("%s != %s", expect.String(), actual.String())
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
    _, err = parser.Parse()
    actual := err.Error()

    if err == nil {
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
