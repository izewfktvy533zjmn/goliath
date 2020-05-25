package scheme

import (
    "fmt"
)

type Number struct {
    Num int
}

func (number *Number) String() string {
    return fmt.Sprintf("%s", number.Num)
}
