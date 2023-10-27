package main

import (
	"strings"
)

var romanNumbers = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var decimalValues = []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

func IsRoman(str string) bool {
	str = strings.ToUpper(str)
	for _, val := range str {
		if val != 'I' && val != 'V' && val != 'X' {
			return false
		}
	}
	return true
}

func ToRoman(num int) string {
	result := ""
	for key, val := range romanNumbers {
		for num >= decimalValues[key] {
			result += val
			num -= decimalValues[key]
		}
	}
	return result
}

func FromRoman(str string) int {
	var result int
	str = strings.ToUpper(str)
	for key, val := range romanNumbers {
		for {
			after, found := strings.CutPrefix(str, val)
			if !found {
				break
			}
			result += decimalValues[key]
			str = after
		}
	}
	return result
}
