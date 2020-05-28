package lexer

import (
    "bufio"
    "errors"
    "strconv"
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

func isWhiteSpace(char byte) bool {
    return char == ' ' || char == '\r' || char == '\n' || char == '\t'
}

func isComment(char byte) bool {
    return char == ';'
}

func isAtomosphere(char byte) bool {
    return isWhiteSpace(char) || isComment(char)
}

func isInitial(char byte) bool {
    return isLetter(char) || isSpecialInitial(char)
}

func isLetter(char byte) bool {
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isSpecialInitial(char byte) bool {
    return char == '!' || char == '$' || char == '%' || char == '&' || char == '*' || char == '/' ||
           char == ':' || char == '<' || char == '=' || char == '>' || char == '?' || char == '^' ||
           char == '_' || char == '~'
}

func isSubsequent(char byte) bool {
    return isInitial(char) || isDigit(char) || isSpecialSubsequent(char)
}

func isDigit(char byte) bool {
    return char >= '0' && char <= '9'
}

func isSpecialSubsequent(char byte) bool {
    return char == '+' || char == '-' || char == '.' || char == '@'
}

func New(in *bufio.Scanner) *Lexer {
    lexer := &Lexer{In: in, NextChar: WHITESPACE_AT_EOL}
    return lexer
}

func (lexer *Lexer)UpdateNextChar() error {
    if lexer.LineIndex == uint64(len(lexer.Line)) {
        lexer.In.Scan()
        text := lexer.In.Text()

        if text == "" {
            return errors.New("EndOfFileException")
        }

        lexer.Line = text + string(WHITESPACE_AT_EOL)
        lexer.NextChar = lexer.Line[0]
        lexer.LineIndex = 1
    } else {
        lexer.NextChar = lexer.Line[lexer.LineIndex]
        lexer.LineIndex++
    }

    return nil
}

func (lexer *Lexer)GetNextToken() (*token.Token, error) {
    char := lexer.NextChar

    for char == WHITESPACE_AT_EOL {
        err := lexer.UpdateNextChar()

        if err != nil {
            return nil, errors.New("SyntaxErrorException")
        }

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

        case '\'':
            lexer.UpdateNextChar()
            return &token.Token{Type: token.QUOTE,  Literal: "'"}, nil

       case '#':
            inputText := ""
            lexer.UpdateNextChar()
            ch := lexer.NextChar

            for isSubsequent(ch) {
                inputText += string(ch)
                err := lexer.UpdateNextChar()

                if err != nil {
                    return nil, errors.New("SyntaxErrorException")
                }

                ch = lexer.NextChar
            }

            if inputText == "t" {
                return &token.Token{Type: token.BOOLEAN, Literal: "#t"}, nil
            } else if inputText == "f" {
                return &token.Token{Type: token.BOOLEAN, Literal: "#f"}, nil
            } else {
                return nil, errors.New("SyntaxErrorException")
            }

        default:
            inputText := ""
            for isSubsequent(char) {
                inputText = inputText + string(char)

                if err := lexer.UpdateNextChar(); err != nil {
                    return nil, errors.New("SyntaxErrorException")
                }

                char = lexer.NextChar
            }

            _, err := strconv.Atoi(inputText)

            if err != nil {
                return &token.Token{Type: token.IDENTIFIER, Literal: inputText}, nil
            } else {
                return &token.Token{Type: token.NUMBER, Literal: inputText}, nil
            }
    }
}
