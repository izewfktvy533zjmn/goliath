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
    lexer := &Lexer{Line: input+string(WHITESPACE_AT_EOL), LineIndex: 0, NextChar: byte(input[0])}
    return lexer
}

func (lexer *Lexer)UpdateNextChar() {
    // TODO: 文字列の読み込み処理と文字列の終端処理
    //if lexer.LineIndex == len(lexer.Line) {


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
            ch := lexer.NextChar
            lexer.UpdateNextChar()

            if lexer.NextChar == WHITESPACE_AT_EOL {
                if ch == 't' {
                    return &token.Token{Type: token.BOOLEAN, Literal: "#t"}
                } else if ch == 'f' {
                    return &token.Token{Type: token.BOOLEAN, Literal: "#f"}
                } else {
                    return &token.Token{Type: token.UNKNOWN, Literal: "unknown"}
                }
            } else {
                return &token.Token{Type: token.UNKNOWN, Literal: "unknown"}
            }

        default:
            return &token.Token{Type: token.UNKNOWN, Literal: "unknown"}
    }
}
