package goLeet

import "fmt"

var romanMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// scores an 84% on runtime and 83.58% on memory on leetcode
func romanToInt(s string) int {
	sum := 0
	fmt.Println(s)
	for i, c := range s {
		minusOperation := false
		lookAhead := ' '

		if len(s) > i + 1 {
			lookAhead = rune(s[i+1])
		}

		val := romanMap[c]
		
		//sum += val
		switch c {
		case 'I':
			if lookAhead == 'V' || lookAhead == 'X' {
				sum -= val
				minusOperation = true
			}
		case 'X':
			if lookAhead == 'L' || lookAhead == 'C' {
				sum -= val
				minusOperation = true
			}
		case 'C':
			if lookAhead == 'D' || lookAhead == 'M' {
				sum -= val
				minusOperation = true
			}
		}
		if !minusOperation {
			sum += val
		}
	}
	return sum
}