package token

import (
    "testing"
    "strconv"
)

func TestString(test *testing.T) {
    token := new(Token)
    token.Type = NUMBER
    token.Literal = "1"

    expect := "Token (Number, 1)"
    actual := token.String()

    if expect != actual {
        test.Errorf("%s != %s", expect, actual)
    }
}

func TestIsTokenType(test *testing.T) {
    token := new(Token)
    token.Type = NUMBER
    token.Literal = "1"

    expect := true
    actual := token.IsTokenType(NUMBER)

    if expect != actual {
        test.Errorf("%s != %s", strconv.FormatBool(expect), strconv.FormatBool(actual))
    }

    expect = false
    actual = token.IsTokenType(BOOLEAN)

    if expect != actual {
        test.Errorf("%s != %s", strconv.FormatBool(expect), strconv.FormatBool(actual))
    }
}

func TestGetValue_Number(test *testing.T) {
    token := new(Token)
    token.Type = NUMBER
    token.Literal = "1"

    expect := 1
    actual, ok := token.GetValue().(int)

    if !ok || expect != actual {
        test.Errorf("%d != %d", expect, actual)
    }
}

func TestGetValue_Boolean_true(test *testing.T) {
    token := new(Token)
    token.Type = BOOLEAN
    token.Literal = "#t"

    expect := true
    actual, ok := token.GetValue().(bool)

    if !ok || expect != actual {
        test.Errorf("%s != %s", strconv.FormatBool(expect), strconv.FormatBool(actual))
    }
}

func TestGetValue_Boolean_false(test *testing.T) {
    token := new(Token)
    token.Type = BOOLEAN
    token.Literal = "#f"

    expect := false
    actual, ok := token.GetValue().(bool)

    if !ok || expect != actual {
        test.Errorf("%s != %s", strconv.FormatBool(expect), strconv.FormatBool(actual))
    }
}
