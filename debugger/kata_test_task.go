package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var intToRoman = []int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

const (
	Negative  = "Выдача паники, так как в римской системе нет отрицательных чисел."
	DiffType  = "Выдача паники, так как используются одновременно разные системы счисления."
	Format    = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	ErrValues = "Выдача паники, так как вводимые значения меньше одного или больше десяти."
	Zero      = "Выдача паники, так как в римской системе счисления нет значения ноль."
)

func overDigTest(str string) {
	reg, _ := regexp.Compile(`[+\-*/]`)
	digits := reg.Split(str, -1)
	if len(digits) > 2 || len(digits) < 2 {
		panic(Format)
	}
}

func overSymbolsTest(signs []string) {
	if len(signs) > 3 || len(signs) < 3 {
		panic(Format)
	}
}

func typecheck(signs []string) int {
	var count int
	_, err1 := strconv.Atoi(signs[0])
	if err1 != nil {
		count++
	}
	_, err2 := strconv.Atoi(signs[2])
	if err2 != nil {
		count++
	}
	return count
}

func result(Val1, Val2 int, operate string) int {
	var res int
	switch operate {
	case "+":
		res = Val1 + Val2
	case "-":
		res = Val1 - Val2
	case "*":
		res = Val1 * Val2
	case "/":
		res = Val1 / Val2
	}
	return res
}

func convIntToRoman(result int) string {
	var romanRes string
	for _, val := range intToRoman {
		for i := val; i <= result; {
			for idx, rValue := range romanMap {
				if rValue == i {
					romanRes += idx
					result -= i
				}
			}
		}
	}
	return romanRes
}

func interpDigType(signs []string, convErr int) {
	switch convErr {
	case 0:
		Val1, _ := strconv.Atoi(signs[0])
		Val2, _ := strconv.Atoi(signs[2])
		if Val1 <= 0 || Val1 >= 11 || Val2 <= 0 || Val2 >= 11 {
			panic(ErrValues)
		} else {
			fmt.Println(result(Val1, Val2, signs[1]))
		}
	case 1:
		panic(DiffType)
	case 2:
		var romInt []int
		for id, elem := range signs {
			if id == 1 {
				continue
			}
			value, check := romanMap[elem]
			if value > 0 && value < 11 && check {
				romInt = append(romInt, value)
			} else {
				panic(ErrValues)
			}
		}
		res := result(romInt[0], romInt[1], signs[1])
		if res < 0 {
			panic(Negative)
		} else if res == 0 {
			panic(Zero)
		} else {
			fmt.Println(convIntToRoman(res))
		}

	}
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSpace(str)
	overDigTest(str)
	signs := strings.Fields(str)
	overSymbolsTest(signs)
	convErr := typecheck(signs)
	interpDigType(signs, convErr)
}
