package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var priority = map[string]int{
	"(": 0,
	")": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if math.Abs(b) < 1e-9 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}
	return 0, errors.New("wrong operator")
}

func makeCalc(numbers, operators *Stack) error {
	if numbers.Size() < 2 || operators.Empty() {
		return errors.New("wrong expression")
	}

	num2 := numbers.Pop().(float64)
	num1 := numbers.Pop().(float64)
	op := operators.Pop().(string)

	res, err := calculate(num1, num2, op)
	if err != nil {
		return err
	}

	numbers.Push(res)
	return nil
}

func Calc(expression string) (float64, error) {
	expr, err := convertString(expression)
	if err != nil {
		return 0, err
	}

	exp := strings.Split(expr, " ")

	var numbers Stack
	var operators Stack

	for _, elem := range exp {
		if num, err := strconv.ParseFloat(elem, 64); err == nil {
			numbers.Push(num)
		} else if elem == "(" {
			operators.Push(elem)
		} else if elem == ")" {
			if operators.Empty() {
				return 0, errors.New("wrong expression")
			}
			for operators.Top() != "(" {
				err := makeCalc(&numbers, &operators)
				if err != nil {
					return 0, err
				}
			}
			operators.Pop()
		} else {
			for !operators.Empty() && priority[elem] <= priority[operators.Top().(string)] {
				err := makeCalc(&numbers, &operators)
				if err != nil {
					return 0, err
				}
			}
			operators.Push(elem)
		}
	}

	for !operators.Empty() {
		err := makeCalc(&numbers, &operators)
		if err != nil {
			return 0, err
		}
	}

	if numbers.Size() != 1 || !operators.Empty() {
		return 0, errors.New("wrong expression")
	}

	res := numbers.Pop().(float64)
	return res, nil
}

func convertString(str string) (string, error) {
	res := ""
	prevNum := false

	for _, elem := range str {
		if elem >= '0' && elem <= '9' {
			if !prevNum {
				res += " "
			}
			res += string(elem)

			prevNum = true
			continue
		}

		prevNum = false

		if _, ok := priority[string(elem)]; ok {
			res += " " + string(elem)
		} else if elem == ' ' {
			continue
		} else {
			return "", errors.New("wrong elem")
		}
	}
	return res[1:], nil
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Use case: go run calc.go \"(1 + 2) - 3\"")
		log.Fatal("wrong input")
		return
	}

	res, err := Calc(os.Args[1])

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res)
}
