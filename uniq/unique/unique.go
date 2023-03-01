package unique

import (
	"fmt"
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

func cutLine(str string, options Options) (res string) {
	if len(str) == 0 {
		return str
	}

	res = ""
	cutStr := strings.Fields(str)[options.SkipFields:]

	for i := range cutStr {
		res = fmt.Sprintf("%s %s", res, cutStr[i])
	}

	if len(res) <= options.SkipChars+1 {
		return ""
	}

	return res[options.SkipChars+1:]
}

func equal(str1, str2 string, options Options) bool {
	if options.Ignore {
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

	var repeats int
	prev := cutLine(str[0], options)
	if (options.Duplicate && repeats == 0) || (options.Unique && repeats != 0) {
		res = append(res, str[0])
	}

	for i := 1; i < len(str); i++ {
		cur := cutLine(str[i], options)

		if equal(prev, cur, options) {
			repeats++
			continue
		}

		if options.Count {
			res = append(res, strconv.Itoa(repeats+1)+" "+str[i-1])
		} else if (options.Duplicate && repeats == 0) || (options.Unique && repeats != 0) {
			res = append(res, str[i])
		}

		repeats = 0
		prev = cur
	}

	if options.Count {
		res = append(res, strconv.Itoa(repeats+1)+" "+str[len(str)-1])
	}

	return
}
