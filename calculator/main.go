package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operate func(int, int) int

type Expression struct {
	X, Y      int
	Operation Operate
}

var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

func PreparingInputSequence(condition string) []string {
	stringArr := []string{}
	conditionArr := strings.Split(condition, "")

	for _, str := range conditionArr {
		if str != " " {
			stringArr = append(stringArr, str)
		}
	}

	return stringArr
}

func Calculator(condition string, ch chan string) {
	stringArr := PreparingInputSequence(condition)

	var data bool = false

	var exp Expression
	var str string

	for _, elem := range stringArr {
		value, err := strconv.Atoi(elem)
		if err == nil {
			if data == false {
				exp.X = value
				str = elem
				data = true
			} else {
				exp.Y = value
				str = str + " " + elem
			}
		} else {
			exp.Operation = operators[elem]
			str = str + " " + elem
		}
	}

	res := exp.Operation(exp.X, exp.Y)
	ch <- str + " = " + strconv.Itoa(res)
	// fmt.Println(str + " = " + strconv.Itoa(res))
}

func main() {
	inputDataStr := [4]string{"2 + 2", "3 + 6", "7 * 7", "9 / 3"}

	ch := make(chan string)

	for _, condition := range inputDataStr {
		go Calculator(condition, ch)
		fmt.Println(<-ch)
	}
	close(ch)
}
