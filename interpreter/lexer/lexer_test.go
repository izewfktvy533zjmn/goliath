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
    fp.Sync()
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

func TestGetNextToken_Leftpar(test *testing.T) {
    inputText := "("
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.LEFTPAR, Literal: "("})
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

func TestGetNextToken_Rightpar(test *testing.T) {
    inputText := ")"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.RIGHTPAR, Literal: ")"})
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

func TestGetNextToken_Dot(test *testing.T) {
    inputText := "."
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.DOT, Literal: "."})
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

func TestGetNextToken_Quote(test *testing.T) {
    inputText := "'"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.QUOTE, Literal: "'"})
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

func TestGetNextToken_Boolean_true(test *testing.T) {
    inputText := "#t"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
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

func TestGetNextToken_Boolean_false(test *testing.T) {
    inputText := "#f"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.BOOLEAN, Literal: "#f"})
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

func TestGetNextToken_Boolean_error1(test *testing.T) {
    inputText := "#test"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)
    _, err = lexer.GetNextToken()

    expect := errors.New("SyntaxErrorException").Error()
    actual := err.Error()

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestGetNextToken_Boolean_error2(test *testing.T) {
    inputText := "#"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)
    _, err = lexer.GetNextToken()

    expect := errors.New("SyntaxErrorException").Error()
    actual := err.Error()

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }

    fp.Close()

    if err := os.Remove("test.scm"); err != nil {
        panic(err)
    }
}

func TestGetNextToken_Number_1(test *testing.T) {
    inputText := "1"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.NUMBER, Literal: "1"})
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

func TestGetNextToken_Number_10(test *testing.T) {
    inputText := "10"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.NUMBER, Literal: "10"})
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

func TestGetNextToken_Identifier(test *testing.T) {
    inputText := "test"
    fp, err := os.OpenFile("test.scm", os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(fp, inputText)
    fp.Sync()
    fp.Close()

    fp, err = os.Open("test.scm")

    if err != nil {
        panic(err)
    }

    in := bufio.NewScanner(fp)
    lexer := New(in)

    expect := *(&token.Token{Type: token.IDENTIFIER, Literal: "test"})
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
