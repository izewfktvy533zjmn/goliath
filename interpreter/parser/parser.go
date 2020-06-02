package parser

import (
    "errors"
    "../token"
    "../lexer"
    "../../scheme/number"
    "../../scheme/boolean"
    "../../scheme/symbol"
    "../../scheme/pair"
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

func (parser *Parser) CheckNestingLevel() error {
    if parser.NestingLevel != 0 {
        tkn, err := parser.Lexer.GetNextToken()

        if err != nil {
            return errors.New("ParseErrorException")
        } else {
            parser.Token = tkn
            return nil
        }
    } else {
        return nil
    }
}

func (parser *Parser) ReadNextToken() error {
    tkn, err := parser.Lexer.GetNextToken()

    if err != nil {
        return errors.New("ParseErrorException")
    } else {
        parser.Token = tkn
        return nil
    }
}

func (parser *Parser) Read() (interface{}, error) {
    switch parser.Token.Type {
        case token.NUMBER:
            value, _ := parser.Token.GetValue().(int)

            if err := parser.CheckNestingLevel(); err != nil {
                return nil, err
            }

            return number.New(value), nil

        case token.BOOLEAN:
            value, _ := parser.Token.GetValue().(bool)

            if err := parser.CheckNestingLevel(); err != nil {
                return nil, err
            }

            return boolean.New(value), nil

        case token.IDENTIFIER:
            value, _ := parser.Token.GetValue().(string)

            if err := parser.CheckNestingLevel(); err != nil {
                return nil, err
            }

            return symbol.New(value), nil

        case token.LEFTPAR:
            parser.NestingLevel++

            if err := parser.ReadNextToken(); err != nil {
                return nil, errors.New("ParseErrorException")
            }

            if parser.Token.Type == token.RIGHTPAR {
                parser.NestingLevel--

                if err := parser.CheckNestingLevel(); err != nil {
                    return nil, err
                }

                return emptylist.New(), nil
            }

            car, err := parser.Read()

            if err != nil {
                return nil, err
            }

            if parser.Token.Type == token.DOT {
                if err := parser.ReadNextToken(); err != nil {
                    return nil, errors.New("ParseErrorException")
                }

                cdr, err := parser.Read()

                if err != nil {
                    return nil, errors.New("ParseErrorException")
                }

                if parser.Token.Type == token.RIGHTPAR {
                    parser.NestingLevel--

                    if err := parser.CheckNestingLevel(); err != nil {
                        return nil, err
                    }

                    ret := pair.New()
                    ret.SetCar(car)
                    ret.SetCdr(cdr)

                    return ret, nil
                } else {
                    return nil, errors.New("ParseErrorException")
                }
            } else {
                if parser.Token.Type == token.RIGHTPAR {
                    parser.NestingLevel--

                    if err := parser.CheckNestingLevel(); err != nil {
                        return nil, err
                    }

                    ret := pair.New()
                    ret.SetCar(car)
                    ret.SetCdr(emptylist.New())

                    return ret, nil
                }

                candidateRet := pair.New()
                candidateRet.SetCar(car)
                candidateRet.SetCdr(emptylist.New())
                var element interface{} = candidateRet

                for {
                    sexp, err := parser.Read()

                    if err != nil {
                        return nil, err
                    }

                    if parser.Token.Type == token.RIGHTPAR {
                        parser.NestingLevel--

                        if err := parser.CheckNestingLevel(); err != nil {
                            return nil, err
                        }

                        candidateSexp := pair.New()
                        candidateSexp.SetCar(sexp)
                        candidateSexp.SetCdr(emptylist.New())

                        element.(*pair.Pair).SetCdr(candidateSexp)
                        ret := candidateRet

                        return ret, nil
                    } else if parser.Token.Type == token.DOT {
                        if err := parser.ReadNextToken(); err != nil {
                            return nil, errors.New("ParseErrorException")
                        }

                        cdr, err := parser.Read()

                        if err != nil || parser.Token.Type != token.RIGHTPAR {
                            return nil, errors.New("ParseErrorException")
                        }

                        parser.NestingLevel--

                        if err := parser.CheckNestingLevel(); err != nil {
                            return nil, err
                        }

                        tmp := pair.New()
                        tmp.SetCar(sexp)
                        tmp.SetCdr(cdr)

                        element.(*pair.Pair).SetCdr(tmp)
                        ret := candidateRet

                        return ret, nil
                    } else {
                        tmp := pair.New()
                        tmp.SetCar(sexp)
                        tmp.SetCdr(emptylist.New())
                        element.(*pair.Pair).SetCdr(tmp)
                        element = element.(*pair.Pair).GetCdr()
                    }
                }

            }

            return nil, errors.New("ParseErrorException")

        default:
            return nil, errors.New("ParseErrorException")
    }
}
