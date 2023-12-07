package main

import (
	"strconv"
	"strings"
)

type GearLocation struct {
	row    int
	column int
	elems  []int
}

func NewGearLocationWith(row, column int) GearLocation {
	return GearLocation{row: row, column: column, elems: make([]int, 0)}
}

func (gl GearLocation) doesExist(el int) bool {
	for i := 0; i < len(gl.elems); i++ {
		if gl.elems[i] == el {
			return true
		}
	}
	return false
}

func GetSomeOfGearRatio(rows []string) int {

	location := make([]GearLocation, 0)

	for i := 0; i < len(rows); i++ {

		switch {
		case i == 0:
			locs, wasFound := GetElementWithGearLocation(0, i, rows[i], rows[i+1])
			if wasFound {
				for i := 0; i < len(locs); i++ {
					AppendLocationSlice(locs[i], &location)
				}
			}
		case i > 0 && i < len(rows)-1:
			locs, wasFound := GetElementWithGearLocation(1, i, rows[i-1], rows[i], rows[i+1])
			if wasFound {
				for i := 0; i < len(locs); i++ {
					AppendLocationSlice(locs[i], &location)
				}
			}
		case i == len(rows)-1:
			locs, wasFound := GetElementWithGearLocation(1, i, rows[i-1], rows[i])
			if wasFound {
				for i := 0; i < len(locs); i++ {
					AppendLocationSlice(locs[i], &location)
				}
			}
		}
	}

	return SumOf(location)
}

func GetElementWithGearLocation(rowToCheck, overallRequiredRow int, rows ...string) ([]GearLocation, bool) {
	requiredRow := rows[rowToCheck]
	var numStr strings.Builder
	locations := make([]GearLocation, 0)
	var isOk = false
	for i := 0; i < len(requiredRow); i++ {
		if isDigit(requiredRow[i]) {
			numStr.WriteByte(requiredRow[i])
			UpdateLocations(&isOk, rowToCheck, overallRequiredRow, i, &locations, rows...)
		}
		if !isDigit(requiredRow[i]) && numStr.Len() > 0 {
			if i < len(requiredRow)-1 && requiredRow[i] != '\r' {
				_, col := Observe(&isOk, rowToCheck, i, ChangeState, isEqual('*', rows[rowToCheck][i]))
				AppendLocationSlice(NewGearLocationWith(overallRequiredRow, col), &locations)
				UpdateLocations(&isOk, rowToCheck, overallRequiredRow, i, &locations, rows...)
			}

			initIndex := i - numStr.Len() - 1

			if initIndex > 0 {

				_, col := Observe(&isOk, rowToCheck, initIndex, ChangeState, isEqual('*', rows[rowToCheck][initIndex]))
				AppendLocationSlice(NewGearLocationWith(overallRequiredRow, col), &locations)
				UpdateLocations(&isOk, rowToCheck, overallRequiredRow, initIndex, &locations, rows...)
			}

			if numStr.Len() > 0 && isOk {

				res, _ := strconv.Atoi(numStr.String())
				locations[len(locations)-1].elems = append(locations[len(locations)-1].elems, res)
			}
			isOk = false
			numStr.Reset()
		}

	}
	return locations, true
}

func UpdateLocations(isOk *bool, rowToCheck, overallRequiredRow, i int, locations *[]GearLocation, rows ...string) {
	if (len(rows) == 2 && rowToCheck == 0) || len(rows) == 3 {
		_, col := Observe(isOk, rowToCheck+1, i, ChangeState, isEqual('*', rows[rowToCheck+1][i]))
		AppendLocationSlice(NewGearLocationWith(overallRequiredRow+1, col), locations)

	}

	if (len(rows) == 2 && rowToCheck == 1) || len(rows) == 3 {
		_, col := Observe(isOk, rowToCheck-1, i, ChangeState, isEqual('*', rows[rowToCheck-1][i]))
		AppendLocationSlice(NewGearLocationWith(overallRequiredRow-1, col), locations)

	}
}

func Observe(target *bool, row, col int, fn func(*bool, bool), res bool) (int, int) {
	before := *target
	fn(target, res)
	if *target && before != *target {
		return row, col
	}
	return -1, -1
}

func doesExist(locs []GearLocation, target GearLocation) (int, bool) {
	for i := 0; i < len(locs); i++ {
		if locs[i].row == target.row && locs[i].column == target.column {
			return i, true
		}
	}
	return -1, false
}

func AppendLocationSlice(location GearLocation, locations *[]GearLocation) {
	if location.row >= 0 && location.column >= 0 {
		if index, exists := doesExist(*locations, location); !exists {
			*locations = append(*locations, location)
		} else {
			for i := 0; i < len(location.elems); i++ {
				(*locations)[index].elems = append((*locations)[index].elems, location.elems[i])
			}
		}
	}
}

func SumOf(locations []GearLocation) int {
	var result int
	for i := 0; i < len(locations); i++ {
		locs := locations[i]
		if len(locs.elems) == 2 {
			tempRes := 1
			for j := 0; j < len(locs.elems); j++ {
				tempRes *= locs.elems[j]
			}
			result += tempRes
		}
	}
	return result
}
