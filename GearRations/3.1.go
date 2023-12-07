package main

import (
	"strconv"
	"strings"
)

func GetSumOfValidNumbers(rows []string) int {
	var result int

	for i := 0; i < len(rows); i++ {

		switch {
		case i == 0:
			result += ParseAdjacentRows(0, rows[i], rows[i+1])
		case i > 0 && i < len(rows)-1:
			result += ParseAdjacentRows(1, rows[i-1], rows[i], rows[i+1])
		case i == len(rows)-1:
			result += ParseAdjacentRows(1, rows[i-1], rows[i])
		}
	}

	return result
}

func ParseAdjacentRows(rowToCheck int, rows ...string) int {
	var result int
	requiredRow := rows[rowToCheck]
	var numStr strings.Builder
	var isOk = false
	for i := 0; i < len(requiredRow); i++ {
		if isDigit(requiredRow[i]) {
			numStr.WriteByte(requiredRow[i])

			if (len(rows) == 2 && rowToCheck == 0) || len(rows) == 3 {
				ChangeState(&isOk, isNotEqual('.', rows[rowToCheck+1][i]))
			}

			if (len(rows) == 2 && rowToCheck == 1) || len(rows) == 3 {
				ChangeState(&isOk, isNotEqual('.', rows[rowToCheck-1][i]))
			}

		}
		if !isDigit(requiredRow[i]) {
			if i < len(requiredRow)-1 && requiredRow[i] != '\r' {

				ChangeState(&isOk, isNotEqual('.', rows[rowToCheck][i]))

				if (len(rows) == 2 && rowToCheck == 0) || len(rows) == 3 {
					ChangeState(&isOk, isNotEqual('.', rows[rowToCheck+1][i]))
				}

				if (len(rows) == 2 && rowToCheck == 1) || len(rows) == 3 {
					ChangeState(&isOk, isNotEqual('.', rows[rowToCheck-1][i]))
				}

			}

			initIndex := i - numStr.Len() - 1

			if initIndex > 0 {

				ChangeState(&isOk, isNotEqual('.', rows[rowToCheck][initIndex]))

				if (len(rows) == 2 && rowToCheck == 0) || len(rows) == 3 {
					ChangeState(&isOk, isNotEqual('.', rows[rowToCheck+1][initIndex]))

				}

				if (len(rows) == 2 && rowToCheck == 1) || len(rows) == 3 {
					ChangeState(&isOk, isNotEqual('.', rows[rowToCheck-1][initIndex]))

				}

			}

			if numStr.Len() > 0 && isOk {

				res, _ := strconv.Atoi(numStr.String())
				result += res
			}
			isOk = false
			numStr.Reset()
		}

	}
	return result
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func isNotEqual(target, b byte) bool {
	return target != b
}

func ChangeState(init *bool, target bool) {
	if !*init {
		*init = target
	}
}

func isEqual(target, b byte) bool {
	return target == b
}
