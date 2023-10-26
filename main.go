package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindOperator(str string) rune {
	for _, val := range str {
		if val == '+' || val == '-' || val == '/' || val == '*' {
			return val
		}
	}
	return '?'
}

func FindOperands(str string, operator rune) (string, string) {
	result := strings.Split(str, string(operator))
	if len(result) != 2 {
		return "", ""
	}
	return result[0], result[1]
}

func IsRim(str string) bool {
	str = strings.ToUpper(str)
	for _, val := range str {
		if val != 'I' && val != 'V' && val != 'X' {
			return false
		}
	}
	return true
}

func Calc(op1, op2 string, op rune) (int, error) {
	num1, err := strconv.Atoi(op1)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.Atoi(op2)
	if err != nil {
		return 0, err
	}
	switch op {
	case '+':
		return num1 + num2, nil
	case '-':
		return num1 - num2, nil
	case '*':
		return num1 * num2, nil
	case '/':
		if num2 == 0 {
			return 0, errors.New("divide by zero")
		} else {
			return num1 / num2, nil
		}
	}
	return 0, errors.New("unknown error")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Write operation: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			continue
		}
		op := FindOperator(str)
		if op == '?' {
			fmt.Printf("error: %s\n", "wrong operator")
			continue
		}
		operand1, operand2 := FindOperands(str, op)
		if operand1 == "" || operand2 == "" {
			fmt.Printf("error: %s\n", "wrong operator")
			continue
		}
		operand2 = strings.Trim(operand2, "\r\n")
		if IsRim(operand1) && IsRim(operand2) {
			continue //TODO
		} else if !IsRim(operand1) && !IsRim(operand2) {
			result, err := Calc(operand1, operand2, op)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				continue
			}
			fmt.Printf("%d\n", result)
		} else {
			fmt.Printf("error: %s\n", "wrong operand")
		}
	}
}
