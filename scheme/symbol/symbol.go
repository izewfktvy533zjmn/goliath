package symbol

type Symbol struct {
    Value string
}

func New(value string) *Symbol {
    return &Symbol{Value: value}
}

func (symbol *Symbol) GetValue() string {
    return symbol.Value
}

func (symbol *Symbol) String() string {
    return symbol.Value
}
