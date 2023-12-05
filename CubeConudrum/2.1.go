package main

import (
	"fmt"
	"strconv"
	"strings"
)

var CubesToMaxAmount = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func GetVerifiedIDSum(rows []string) int {
	var result int

	for _, row := range rows {

		id := GetGameID(row)

		row = row[strings.Index(row, ":")+1:]
		stats := ParseString(row)
		isOk := true
		for k, v := range CubesToMaxAmount {
			for i := 0; i < len(stats); i++ {
				if stats[i][k] > v {
					isOk = false
				}
			}

		}

		if isOk {
			result += id
		} else {
			fmt.Println(id)
		}
	}

	return result
}

func ParseString(line string) []map[string]int {
	stats := make([]map[string]int, 0)
	subsets := make([]string, 0)
	subsets = strings.Split(line, "; ")

	for j, sub := range subsets {
		stats = append(stats, make(map[string]int))
		var cube strings.Builder
		var numStr strings.Builder
		for i := 0; i < len(sub); i++ {
			if isDigit(sub[i]) {
				numStr.WriteByte(sub[i])
			}
			if isLetter(sub[i]) {
				cube.WriteByte(sub[i])
			}
			if isDelimiter(',', sub[i]) || i == len(sub)-1 {
				num, _ := strconv.Atoi(numStr.String())
				stats[j][cube.String()] += num
				cube.Reset()
				numStr.Reset()
			}

		}

	}

	for i := 0; i < len(stats); i++ {
		for k, v := range stats[i] {
			fmt.Println(k, v)
		}
		fmt.Println("====")
	}

	return stats
}

func GetGameID(line string) int {
	i := 0
	var ID int
	var sb strings.Builder
	for line[i] != ':' {
		if isDigit(line[i]) {
			sb.WriteByte(line[i])
		}
		i++
	}
	ID, _ = strconv.Atoi(sb.String())

	return ID
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func isLetter(b byte) bool {
	return b >= 97 && b <= 122
}

func isDelimiter(delim, b byte) bool {
	return b == delim
}
