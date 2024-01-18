package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanArray = []string{
	"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Совершите математическую операцию (+, -, /, *) римскими или арабскими цифрами: ")
	scanner.Scan()
	expression := scanner.Text()

	result, err := parse(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(result)
}

func parse(expression string) (string, error) {
	var num1, num2 int
	var oper, result string
	var isRoman bool

	operands := regexp.MustCompile(`\s*[+\-*/]\s*`).Split(expression, -1)

	if len(operands) != 2 {
		return "", fmt.Errorf("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *). операцией")
	}

	oper = detectOperation(expression)
	if oper == "" {
		return "", fmt.Errorf("неподдерживаемая математическая операция")
	}

	if isRomanNumber(operands[0]) && isRomanNumber(operands[1]) {
		num1 = convertToArabian(operands[0])
		num2 = convertToArabian(operands[1])
		isRoman = true
	} else if !isRomanNumber(operands[0]) && !isRomanNumber(operands[1]) {
		var err1, err2 error
		num1, err1 = strconv.Atoi(operands[0])
		num2, err2 = strconv.Atoi(operands[1])

		if err1 != nil || err2 != nil {
			return "", fmt.Errorf("конвертации строки в число")
		}

		isRoman = false
	} else {
		return "", fmt.Errorf("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *). операцией")
	}

	if num1 > 10 || num2 > 10 {
		return "", fmt.Errorf("числа должны быть от 1 до 10")
	}

	arabian := calc(num1, num2, oper)

	if isRoman {
		if arabian <= 0 {
			return "", fmt.Errorf("в римской системе нет отрицательных чисел")
		}
		result = convertToRoman(arabian)
	} else {
		result = strconv.Itoa(arabian)
	}

	return result, nil
}

func detectOperation(expression string) string {
	switch {
	case strings.Contains(expression, "+"):
		return "+"
	case strings.Contains(expression, "-"):
		return "-"
	case strings.Contains(expression, "*"):
		return "*"
	case strings.Contains(expression, "/"):
		return "/"
	default:
		return ""
	}
}

func calc(a, b int, oper string) int {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func isRomanNumber(val string) bool {
	for _, roman := range romanArray {
		if val == roman {
			return true
		}
	}
	return false
}

func convertToArabian(roman string) int {
	for i, val := range romanArray {
		if roman == val {
			return i
		}
	}
	return -1
}

func convertToRoman(arabian int) string {
	return romanArray[arabian]
}
