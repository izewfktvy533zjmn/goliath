package boolean

type Boolean struct {
    Value bool
}

func New(value bool) *Boolean {
    return &Boolean{Value: value}
}

func (boolean *Boolean) GetValue() bool {
    return boolean.Value
}

func (boolean *Boolean) String() string {
    if boolean.Value {
        return "#t"
    } else {
        return "#f"
    }
}
