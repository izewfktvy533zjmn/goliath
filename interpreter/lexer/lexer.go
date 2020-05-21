package lexer

import (
    //"errors"
    "../token"
)

type Lexer struct {
    Line      string
    LineIndex uint64
    NextChar  byte
}

const (
    WHITESPACE_AT_EOL = ' ';
)

func New(input string) *Lexer {
    lexer := &Lexer{Line: input, LineIndex: 0, NextChar: byte(input[0])}
    return lexer
}

func (lexer *Lexer)UpdateNextChar() {
    lexer.LineIndex++
    lexer.NextChar = lexer.Line[lexer.LineIndex]
}

func (lexer *Lexer)GetNextToken() *token.Token {
    char := lexer.NextChar

    for char == WHITESPACE_AT_EOL {
        lexer.UpdateNextChar()
        char = lexer.NextChar
    }

    switch char {
        case '(':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.LEFTPAR, Literal: "("}

        case ')':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.RIGHTPAR, Literal: ")"}

        case '.':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.DOT, Literal: "."}

        case '#':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.BOOLEAN, Literal: "true"}
            return &token.Token{Type: token.BOOLEAN, Literal: "false"}

        default:
            //return errors.New("test")
            return &token.Token{Type: token.NUMBER, Literal: "1"}
    }
}
