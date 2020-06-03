package pair

import (
    "fmt"
    "../emptylist"
    "../../scheme"
)

type Pair struct {
    Car interface{}
    Cdr interface{}
}

func New() *Pair {
    return &Pair{}
}

func (pair *Pair) SetCar(car interface{}) {
    pair.Car = car
}

func (pair *Pair) SetCdr(cdr interface{}) {
    pair.Cdr = cdr
}

func (pair *Pair) GetCar() interface{} {
    return pair.Car
}

func (pair *Pair) GetCdr() interface{} {
    return pair.Cdr
}

func (pair *Pair) ToString() string {
    if _, ok := pair.Cdr.(*emptylist.EmptyList); ok {
        return fmt.Sprintf("(%s)", pair.GetCar().(scheme.SExp).ToString())
    } else if _, ok := pair.Cdr.(*Pair); !ok {
        return fmt.Sprintf("(%s . %s)", pair.GetCar().(scheme.SExp).ToString(), pair.GetCdr().(scheme.SExp).ToString())
    } else {
        tmpCar := pair.GetCdr().(*Pair).GetCar()
        tmpCdr := pair.GetCdr().(*Pair).GetCdr()

        str := "" + pair.GetCar().(scheme.SExp).ToString()

        for {
            if _, ok := tmpCar.(*emptylist.EmptyList); ok {
                return "(" + str + ")"
            } else if _, ok := tmpCdr.(*emptylist.EmptyList); ok {
                return "(" + str + " " + tmpCar.(scheme.SExp).ToString() + ")"
            } else if _, ok := tmpCdr.(*Pair); !ok {
                return "(" + str + " " + tmpCar.(scheme.SExp).ToString() + " . " + tmpCdr.(scheme.SExp).ToString() + ")"
            } else {
                str += " " + tmpCar.(scheme.SExp).ToString()
                tmpCar = tmpCdr.(*Pair).GetCar()
                tmpCdr = tmpCdr.(*Pair).GetCdr()
            }
        }
    }
}
