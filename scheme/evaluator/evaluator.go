package evaluator

import (
    "../../scheme"
    "../symbol"
    "../frame"
)

type Evaluator struct {
    Frame *frame.Frame
}

func New(frame *frame.Frame) *Evaluator {
    return &Evaluator{Frame: frame}
}

func (evaluator *Evaluator) Evaluate(sexp interface {}) (scheme.SExp, error) {
    switch sexp := sexp.(type) {
        case *symbol.Symbol:
            if ret, err := evaluator.Frame.Refer(sexp); err != nil {
                return nil, err
            } else {
                return ret.(scheme.SExp), err
            }

        default:
            return sexp.(scheme.SExp), nil
    }
}
