package lexer

import (
    "errors"
    "bufio"
    "../token"
)

type Lexer struct {
    In        *bufio.Scanner
    Line      string
    LineIndex uint64
    NextChar  byte
}

const (
    WHITESPACE_AT_EOL = ' ';
)

func New(in *bufio.Scanner) *Lexer {
    lexer := &Lexer{In: in, NextChar: WHITESPACE_AT_EOL}
    return lexer
}

func (lexer *Lexer)UpdateNextChar() error {
    if lexer.LineIndex == uint64(len(lexer.Line)) {
        text := lexer.In.Text()
        if text == "" {
            lexer.In.Scan()
            text = lexer.In.Text()

            if text == "" {
                return errors.New("EndOfFileException")
            }
        }

        lexer.Line = text + string(WHITESPACE_AT_EOL)
        lexer.LineIndex = 0;
        lexer.NextChar = lexer.Line[lexer.LineIndex]
    } else {
        lexer.LineIndex++
        lexer.NextChar = lexer.Line[lexer.LineIndex]
    }

    return nil
}

func (lexer *Lexer)GetNextToken() (*token.Token, error) {
    char := lexer.NextChar

    for char == WHITESPACE_AT_EOL {
        lexer.UpdateNextChar()
        char = lexer.NextChar
    }

    switch char {
        case '(':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.LEFTPAR, Literal: "("}, nil

        case ')':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.RIGHTPAR, Literal: ")"}, nil

        case '.':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.DOT, Literal: "."}, nil

        case '#':
            lexer.UpdateNextChar()
            ch := lexer.NextChar
            lexer.UpdateNextChar()

            if lexer.NextChar == WHITESPACE_AT_EOL {
                if ch == 't' {
                    return &token.Token{Type: token.BOOLEAN, Literal: "#t"}, nil
                } else if ch == 'f' {
                    return &token.Token{Type: token.BOOLEAN, Literal: "#f"}, nil
                } else {
                    return nil, errors.New("SyntaxErrorException")
                }
            } else {
                return nil, errors.New("SyntaxErrorException")
            }

        default:
            return nil, errors.New("SyntaxErrorException")

    }
}
