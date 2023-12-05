package main

import (
	"fmt"
	"strconv"
	"strings"
)

const defValue = -1

var wordsToDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getDigitsSum(str []string) int {
	var result int

	for _, s := range str {
		var first, last = defValue, defValue
		var tempStr strings.Builder
		for i := 0; i < len(s); i++ {
			if isDigit(s[i]) {
				if first == defValue {
					first, _ = strconv.Atoi(string(s[i]))
				}
				last, _ = strconv.Atoi(string(s[i]))
				tempStr.Reset()
				continue
			}

			tempStr.WriteByte(s[i])
			for k, v := range wordsToDigits {
				if strings.Contains(tempStr.String(), k) {
					fmt.Println(tempStr.String())
					if first == defValue {
						first = v
					}
					last = v
					tempStr.Reset()
					tempStr.WriteByte(s[i])
					break
				}
			}
		}
		result += 10*first + last
	}

	return result
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}
