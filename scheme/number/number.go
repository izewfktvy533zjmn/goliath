package number

import (
    "fmt"
)

type Number struct {
    Value int
}

func New(value int) *Number {
    return &Number{Value: value}
}

func (number *Number) GetValue() int {
    return number.Value
}

func (number *Number) String() string {
    return fmt.Sprintf("%d", number.Value)
}
