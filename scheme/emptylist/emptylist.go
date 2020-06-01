package emptylist

import (
    "fmt"
)

type EmptyList struct {
    Value *EmptyList
}

func New() *EmptyList {
    return &EmptyList{}
}

func (emptylist *EmptyList) String() string {
    return fmt.Sprintf("()")
}
