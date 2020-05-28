package identifier

type Identifier struct {
    Value string
}

func New(value string) *Identifier {
    return &Identifier{Value: value}
}

func (identifier *Identifier) GetValue() string {
    return identifier.Value
}

func (identifier *Identifier) String() string {
    return identifier.Value
}
