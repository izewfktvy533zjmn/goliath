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
    NUMBER       = "Number"
    BOOLEAN      = "Boolean"
    INDENTIFIER  = "Indentifier"
    LEFTPAR      = "Leftpar"
    RIGHTPAR     = "Rightpar"
    DOT          = "Dot"
    QUOTE        = "Quote"
    UNKNOWN      = "Unknown"
)

func (token *Token) String() string {
    switch token.Type {
        case NUMBER, BOOLEAN, INDENTIFIER:
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
            ret, _ := strconv.ParseBool(token.Literal)
            return ret

        default:
            return token.Literal
    }
}
