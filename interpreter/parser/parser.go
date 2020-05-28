package parser

import (
    "errors"
    "../token"
    "../lexer"
    "../../scheme/number"
    "../../scheme/boolean"
    "../../scheme/identifier"
)

type Parser struct {
    Lexer        *lexer.Lexer
    Token        *token.Token
    NestingLevel uint64
}

func New(l *lexer.Lexer) *Parser {
    return &Parser{Lexer: l, Token: nil, NestingLevel: 0}
}

func (parser *Parser) Parse() (interface{}, error) {
    parser.NestingLevel = 0
    token, err := parser.Lexer.GetNextToken()
    parser.Token = token

    if err != nil {
        return nil, errors.New("ParseErrorException")
    }

    return parser.Read()
}

func (parser *Parser) Read() (interface{}, error) {
    switch parser.Token.Type {
        case token.NUMBER:
            value, _ := parser.Token.GetValue().(int)
            return number.New(value), nil

        case token.BOOLEAN:
            value, _ := parser.Token.GetValue().(bool)
            return boolean.New(value), nil

        case token.IDENTIFIER:
            value, _ := parser.Token.GetValue().(string)
            return identifier.New(value), nil

        default:
            return nil, errors.New("ParseErrorException")
    }
}
