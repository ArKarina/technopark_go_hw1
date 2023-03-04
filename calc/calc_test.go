package main

import (
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
	t.Parallel()
	for _, test := range tests {
		test := test
		t.Run(test.expression, func(t *testing.T) {
			t.Parallel()
			result, err := Calc(test.expression)
			require.NoError(t, err)
			require.Equal(t, test.res, result, "Check Calc func")
		})
	}
}

var testsBad = []struct {
	expression string
	err        string
}{
	{"((1+2)-3)))",
		"wrong expression"},
	{"( 1 +   2) * 3  !  ",
		"wrong elem"},
	{"1+2* (3+4/0-(1+2))*2+1",
		"division by zero"},
	{"( 1 + 2 4  2) * 3  !  ",
		"wrong elem"},
}

func TestCalcIncorrect(t *testing.T) {
	t.Parallel()
	for _, test := range testsBad {
		test := test
		t.Run(test.expression, func(t *testing.T) {
			t.Parallel()
			_, err := Calc(test.expression)
			require.ErrorContains(t, err, test.err, "Check Calc func")
		})
	}
}
