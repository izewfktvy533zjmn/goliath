package pair

import (
    "fmt"
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

func (pair *Pair) String() string {
    return fmt.Sprintf("(pair)")
}
