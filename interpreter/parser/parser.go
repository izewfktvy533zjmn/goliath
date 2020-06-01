package parser

import (
    "errors"
    "../token"
    "../lexer"
    "../../scheme/number"
    "../../scheme/boolean"
    "../../scheme/identifier"
    //"../../scheme/pair"
    "../../scheme/emptylist"
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
    tkn, err := parser.Lexer.GetNextToken()
    parser.Token = tkn

    if err != nil {
        return nil, errors.New("ParseErrorException")
    }

    return parser.Read()
}

func (parser *Parser) Read() (interface{}, error) {
    switch parser.Token.Type {
        case token.NUMBER:
            value, _ := parser.Token.GetValue().(int)

            if parser.NestingLevel != 0 {
                tkn, err := parser.Lexer.GetNextToken()

                if err != nil {
                    return nil, errors.New("ParseErrorException")
                }

                parser.Token = tkn
            }

            return number.New(value), nil

        case token.BOOLEAN:
            value, _ := parser.Token.GetValue().(bool)

            if parser.NestingLevel != 0 {
                tkn, err := parser.Lexer.GetNextToken()

                if err != nil {
                    return nil, errors.New("ParseErrorException")
                }

                parser.Token = tkn
            }

            return boolean.New(value), nil

        case token.IDENTIFIER:
            value, _ := parser.Token.GetValue().(string)

            if parser.NestingLevel != 0 {
                tkn, err := parser.Lexer.GetNextToken()

                if err != nil {
                    return nil, errors.New("ParseErrorException")
                }

                parser.Token = tkn
            }

            return identifier.New(value), nil

        case token.LEFTPAR:
            parser.NestingLevel++
            tkn, err := parser.Lexer.GetNextToken()

            if err != nil {
                return nil, errors.New("ParseErrorException")
            }

            parser.Token = tkn

            if parser.Token.Type == token.RIGHTPAR {
                parser.NestingLevel--

                if parser.NestingLevel != 0 {
                    tkn, err := parser.Lexer.GetNextToken()

                    if err != nil {
                        return nil, errors.New("ParseErrorException")
                    }

                    parser.Token = tkn
                }

                return emptylist.New(), nil
            }
            return nil, errors.New("ParseErrorException")


        default:
            return nil, errors.New("ParseErrorException")
    }
}
