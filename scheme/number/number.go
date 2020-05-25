package number

import (
    "fmt"
)

type Number struct {
    Num int
}

func New(num int) *Number {
    return &Number{Num: num}
}

func (number *Number) GetValue() int {
    return number.Num
}

func (number *Number) String() string {
    return fmt.Sprintf("%s", number.Num)
}
