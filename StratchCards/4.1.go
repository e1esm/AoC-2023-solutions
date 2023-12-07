package main

import (
	"strings"
)

func GetOverallPointsOfCards(rows []string) int {
	var result int

	for i := 0; i < len(rows); i++ {
		result += ParseCardStats(rows[i])
	}

	return result
}

func ParseCardStats(stat string) int {

	pointSl := strings.Split(strings.ReplaceAll(stat[strings.Index(stat, ":")+1:len(stat)-1], "  ", " "), "| ")
	pointSl[0] = strings.Trim(pointSl[0], " ")

	winningNumbers := strings.Split(pointSl[0], " ")
	obtainedNumbers := strings.Split(pointSl[1], " ")
	if len(winningNumbers) > len(obtainedNumbers) {
		return FindPointsOf(winningNumbers, obtainedNumbers)
	}

	return FindPointsOf(obtainedNumbers, winningNumbers)
}

func FindPointsOf(target, current []string) int {
	var points int
	matches := 0
	for i := 0; i < len(target); i++ {
		for j := 0; j < len(current); j++ {
			if target[i] == current[j] {
				matches++
				if matches == 1 {
					points++
					continue
				}
				points *= 2

			}
		}
	}
	return points
}
