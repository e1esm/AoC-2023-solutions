package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetPowerOfFewestValidCubes(rows []string) int {
	var result int

	for _, row := range rows {
		row = row[strings.Index(row, ":")+1:]
		result += PowerOf(GetFewestCubesAmount(row))
	}

	return result
}

func GetFewestCubesAmount(row string) []int {
	stats := strings.Split(row, ";")
	cubesToIndex := map[string]int{"red": 0, "blue": 1, "green": 2}
	arr := make([]int, 3)
	for i := 0; i < len(stats); i++ {
		var numStr strings.Builder
		var cube strings.Builder
		for j := 0; j < len(stats[i]); j++ {
			b := stats[i][j]
			if isDigit(b) {
				numStr.WriteByte(b)
			}
			if isLetter(b) {
				cube.WriteByte(b)
			}
			if isDelimiter(stats[i][j], ',') || j == len(stats[i])-1 {
				num, _ := strconv.Atoi(numStr.String())
				if arr[cubesToIndex[cube.String()]] < num {
					arr[cubesToIndex[cube.String()]] = num
				}
				numStr.Reset()
				cube.Reset()
			}
		}

	}
	return arr
}

func PowerOf(arr []int) int {
	result := 1
	for i := 0; i < len(arr); i++ {
		result *= arr[i]
		fmt.Println(arr[i])
	}
	fmt.Println("====")
	return result
}
