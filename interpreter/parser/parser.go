package parser

import (
    "errors"
    "../token"
    "../lexer"
    "../../scheme"
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
    parser.Token, _ = parser.Lexer.GetNextToken()

    return parser.Read()
}

func (parser *Parser) Read() (interface{}, error) {
    switch parser.Token.Type {
        case token.NUMBER:
            val, _ := parser.Token.GetValue().(int)
            return &scheme.Number{Num: val}, nil

        default:
            return nil, errors.New("ParseErrorException")
    }
}
