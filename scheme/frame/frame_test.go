package frame

import (
    "testing"
    "../number"
    "../symbol"

)

func TestBind(test *testing.T) {
    frame := New(nil)

    expect := *number.New(1)
    frame.Bind(symbol.New("test"), number.New(1))

    actual := *frame.Table[*symbol.New("test")].(*number.Number)

    if expect != actual {
        test.Errorf("%v != %v", expect, actual)
    }
}

func TestRefer(test *testing.T) {
    frame := New(nil)

    expect := *number.New(1)
    frame.Bind(symbol.New("test"), number.New(1))
    tmp, err := frame.Refer(symbol.New("test"))

    if err != nil {
        test.Errorf("not find symbol")
        return
    }

    actual := *tmp.(*number.Number)

    if expect != actual {
        test.Errorf("%v != %v", expect, actual)
    }
}
