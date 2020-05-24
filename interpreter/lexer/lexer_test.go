package lexer

import (
    "bufio"
    "errors"
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
    inputText := "test" + string(WHITESPACE_AT_EOL)

    in := bufio.NewScanner(os.Stdin)
    lexer := New(in)
    lexer.Line = inputText
    lexer.LineIndex = 0
    lexer.NextChar = lexer.Line[0]

    lexer.UpdateNextChar()

    var expect byte
    expect = 'e'
    actual := lexer.NextChar

    if expect != actual {
        test.Errorf("%s != %s", string(expect), string(actual))
    }
}

func TestUpdateNextChar_error(test *testing.T) {
    in := bufio.NewScanner(os.Stdin)
    lexer := New(in)

    expect := errors.New("EndOfFileException").Error()
    actual := lexer.UpdateNextChar().Error()

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }
}

func TestGetNextToken(test *testing.T) {
    inputText := "#t" + string(WHITESPACE_AT_EOL)

    in := bufio.NewScanner(os.Stdin)
    lexer := New(in)
    lexer.Line = inputText
    lexer.LineIndex = 0
    lexer.NextChar = lexer.Line[0]

    expect := *(&token.Token{Type: token.BOOLEAN, Literal: "#t"})
    tmp, _ := lexer.GetNextToken()
    actual := *(tmp)

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }
}
