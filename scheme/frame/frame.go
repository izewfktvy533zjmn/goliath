package frame

import (
    "errors"
    "../symbol"
)

type Frame struct {
    Parent *Frame
    Table map[symbol.Symbol]interface{}
}

func New(parent *Frame) *Frame {
    return &Frame{Parent: parent, Table: make(map[symbol.Symbol]interface{})}
}

func (frame *Frame) SearchFrame(sbl *symbol.Symbol) (*Frame, error) {
    for f := frame; f != nil; f = f.Parent {
        if _, ok := f.Table[*sbl]; ok {
            return f, nil
        }
    }

    return nil, errors.New("UnboundVariableException")
}

func (frame *Frame) Refer(sbl *symbol.Symbol) (interface{}, error) {
    if f, err := frame.SearchFrame(sbl); err != nil {
        return nil, err
    } else {
        return f.Table[*sbl], nil
    }
}

func (frame *Frame) Bind(sbl *symbol.Symbol, sexp interface{}) *symbol.Symbol {
    frame.Table[*sbl] = sexp
    return sbl
}

func (frame *Frame) Rebind(sbl *symbol.Symbol, sexp interface{}) (*symbol.Symbol, error) {
    if _, err := frame.Refer(sbl); err != nil {
        return nil, err
    } else {
        return frame.Bind(sbl, sexp), nil
    }
}
