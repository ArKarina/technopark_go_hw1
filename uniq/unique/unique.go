package unique

import (
	"strconv"
	"strings"
)

const Info = "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

type Options struct {
	Count      bool
	Duplicate  bool
	Unique     bool
	SkipFields int
	SkipChars  int
	Ignore     bool
}

func (options Options) Correct() bool {
	if options.Count && options.Duplicate {
		return false
	}
	if options.Count && options.Unique {
		return false
	}
	if options.Duplicate && options.Unique {
		return false
	}

	return true
}

func cutLine(str string, skipFields, skipChars int) (res string) {
	if len(str) == 0 {
		return str
	}

	if len(strings.Fields(str)) <= skipFields {
		return ""
	}

	cutStr := strings.Fields(str)[skipFields:]

	myslice := []string{res, strings.Join(cutStr, "")}
	res = strings.Join(myslice, "")

	if len(res) <= skipChars {
		return ""
	}

	return res[skipChars:]
}

func equal(str1, str2 string, ignore bool) bool {
	if ignore {
		return strings.EqualFold(str1, str2)
	}

	return str1 == str2
}

func Unique(str []string, options Options) (res []string) {
	if len(str) == 0 {
		return
	}

	if !options.Count && !options.Duplicate && !options.Unique {
		options.Unique = true
		options.Duplicate = true
	}

	str = append(str, "\n")
	var repeats int
	prev := cutLine(str[0], options.SkipFields, options.SkipChars)

	for i := 1; i < len(str); i++ {
		cur := cutLine(str[i], options.SkipFields, options.SkipChars)

		if equal(prev, cur, options.Ignore) {
			repeats++
		} else {
			if options.Count {
				res = append(res, strconv.Itoa(repeats+1)+" "+str[i-1])
				repeats = 0
			} else {
				if (options.Duplicate && repeats != 0) || (options.Unique && repeats == 0) {
					res = append(res, str[i-1])
				}
				repeats = 0
			}
		}
		prev = cur
	}

	return res
}
