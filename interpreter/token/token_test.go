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

func TestGetValue(test *testing.T) {
    token_number := new(Token)
    token_number.Type = NUMBER
    token_number.Literal = "1"

    expect := 1
    actual, ok := token_number.GetValue().(int)

    if !ok || expect != actual {
        test.Errorf("%d != %d", expect, actual)
    }

    token_boolean := new(Token)
    token_boolean.Type = BOOLEAN
    token_boolean.Literal = "true"

    expect_b := true
    actual_b, ok := token_boolean.GetValue().(bool)

    if !ok || expect_b != actual_b {
        test.Errorf("%s != %s", strconv.FormatBool(expect_b), strconv.FormatBool(actual_b))
    }
}
