package lexer

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "testing"
    "../token"
)

func TestNew(test *testing.T) {
    in := bufio.NewScanner(os.Stdin)
    lexer := New(in)

    expect := in
    actual := lexer.In

    if expect != actual {
        test.Errorf("%p != %p", expect, actual)
    }
}

func TestUpdateNextChar(test *testing.T) {
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
    lexer := New(in)
    lexer.UpdateNextChar()

    var expect byte
    expect = 't'
    actual := lexer.NextChar

    if expect != actual {
        test.Errorf("%s != %s", string(expect), string(actual))
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestUpdateNextChar_error(test *testing.T) {
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err)
    }
    fp.Close()

    fp, err = os.Open("test.scm")
    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := errors.New("EndOfFileException").Error()
    actual := lexer.UpdateNextChar().Error()

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestGetNextToken(test *testing.T) {
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
    lexer := New(in)

    expect := *(&token.Token{Type: token.BOOLEAN, Literal: "#t"})
    tmp, _ := lexer.GetNextToken()

    actual := *(tmp)

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}
