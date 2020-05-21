package lexer

import (
    "testing"
    "../token"
)

func TestNew_line(test *testing.T) {
    input := "test"

    lexer := New(input)

    expect := input
    actual := lexer.Line

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }
}

func TestNew_lineIndex(test *testing.T) {
    input := "test"

    lexer := New(input)

    var expect uint64
    expect = 0
    actual := lexer.LineIndex

    if expect != actual {
        test.Errorf("%d != %d", expect, actual)
    }
}

func TestNew_nextChar(test *testing.T) {
    input := "test"

    lexer := New(input)

    var expect byte
    expect = 't'
    actual := lexer.NextChar

    if expect != actual {
        test.Errorf("%b != %b", expect, actual)
    }
}

func TestUpdateNextChar(test *testing.T) {
    input := "test"

    lexer := New(input)
    lexer.UpdateNextChar()

    var expect byte
    expect = 'e'
    actual := lexer.NextChar

    if expect != actual {
        test.Errorf("%b != %b", expect, actual)
    }
}

func TestGetNextToken(test *testing.T) {
    input := "(+ 1 1)"

    lexer := New(input)

    expect := *(&token.Token{Type: token.LEFTPAR, Literal: "("})
    actual := *(lexer.GetNextToken())

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }
}
