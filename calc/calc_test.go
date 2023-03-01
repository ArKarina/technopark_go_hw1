package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	expression string
	res        float64
}{
	{"(1+2)-3",
		0},
	{"( 1 +   2) * 3   ",
		9},
	{"1+2* (3+4/2-(1+2))*2+1",
		10},
	{"9+909+9*9",
		999},
}

func TestCalcCorrect(t *testing.T) {
	for _, e := range tests {
		result, _ := Calc(e.expression)
		require.Equal(t, e.res, result, "Check Calc func")
	}
}

var testsBad = []struct {
	expression string
	err        error
}{
	{"((1+2)-3)))",
		errors.New("wrong expression")},
	{"( 1 +   2) * 3  !  ",
		errors.New("wrong elem")},
	{"1+2* (3+4/0-(1+2))*2+1",
		errors.New("division by zero")},
	{"( 1 + 2 4  2) * 3  !  ",
		errors.New("wrong elem")},
}

func TestCalcIncorrect(t *testing.T) {
	for _, e := range testsBad {
		_, err := Calc(e.expression)
		require.Equal(t, e.err, err, "Check Calc func")
	}
}
