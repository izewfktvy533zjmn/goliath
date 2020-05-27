package token

import (
    "fmt"
    "strconv"
)

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    NUMBER     = "Number"
    BOOLEAN    = "Boolean"
    IDENTIFIER = "Identifier"
    LEFTPAR    = "Leftpar"
    RIGHTPAR   = "Rightpar"
    DOT        = "Dot"
    QUOTE      = "Quote"
)

func (token *Token) String() string {
    switch token.Type {
        case NUMBER, BOOLEAN, IDENTIFIER:
            return fmt.Sprintf("Token (%s, %s)", token.Type, token.Literal)

        default:
            return fmt.Sprintf("Token (%s)", token.Type)
    }
}

func (token *Token) IsTokenType(tokenType TokenType) bool {
    if token.Type == tokenType {
        return true
    } else {
        return false
    }
}

func (token *Token) GetValue() interface{} {
    switch token.Type {
        case NUMBER:
            ret, _ := strconv.Atoi(token.Literal)
            return ret

        case BOOLEAN:
            if token.Literal == "#t" {
                return true
            } else {
                return false
            }

        default:
            return token.Literal
    }
}
