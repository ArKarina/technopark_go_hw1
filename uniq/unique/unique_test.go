package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	name    string
	str     []string
	options Options
	res     []string
}{
	{"1",
		[]string{"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
		Options{
			Count:      false,
			Duplicate:  false,
			Unique:     false,
			SkipFields: 0,
			SkipChars:  0,
			Ignore:     false,
		},
		[]string{"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik."}},

	{"2",
		[]string{"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
		Options{
			Count:      true,
			Duplicate:  false,
			Unique:     false,
			SkipFields: 0,
			SkipChars:  0,
			Ignore:     false,
		},
		[]string{"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik."}},

	{"3",
		[]string{"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
		Options{
			Count:      false,
			Duplicate:  true,
			Unique:     false,
			SkipFields: 0,
			SkipChars:  0,
			Ignore:     false,
		},
		[]string{"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik."}},

	{"4",
		[]string{"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
		Options{
			Count:      false,
			Duplicate:  false,
			Unique:     true,
			SkipFields: 0,
			SkipChars:  0,
			Ignore:     false,
		},
		[]string{"",
			"Thanks."}},

	{"5",
		[]string{"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			"",
			"I love MuSIC of Kartik.",
			"I love music of kartik.",
			"Thanks.",
			"I love music of kartik.",
			"I love MuSIC of Kartik."},
		Options{
			Count:      false,
			Duplicate:  false,
			Unique:     false,
			SkipFields: 0,
			SkipChars:  0,
			Ignore:     true,
		},
		[]string{"I LoVe MuSiC.",
			"",
			"I love music of kartik.",
			"Thanks.",
			"I love MuSIC of Kartik."}},

	{"6",
		[]string{"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thank you."},
		Options{
			Count:      false,
			Duplicate:  false,
			Unique:     false,
			SkipFields: 1,
			SkipChars:  0,
			Ignore:     false,
		},
		[]string{"They love music.",
			"",
			"We love music of Kartik.",
			"Thank you."}},

	{"7",
		[]string{"I love music.",
			"A love music.",
			"I love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks."},
		Options{
			Count:      false,
			Duplicate:  false,
			Unique:     false,
			SkipFields: 0,
			SkipChars:  1,
			Ignore:     false,
		},
		[]string{"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks."}},
}

func TestTUniq(t *testing.T) {
	t.Parallel()
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, test.res, Unique(test.str, test.options), "Check Unique func")
		})
	}
}
